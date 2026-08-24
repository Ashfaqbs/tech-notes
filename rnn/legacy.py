import re
import torch
import torch.nn as nn
import torch.optim as optim
torch.manual_seed(0)

# ---------- tiny training corpus ----------
sentences = [
    "i like machine learning",
    "i like pizza",
    "i like coding in python",
    "rnn models process sequences",
    "transformers use attention",
    "deep lea   rning is powerful",
    "pytorch makes building models easy",
    "we predict the next word",
    "this is a simple dataset",
    "neural networks learn patterns"
]
sentences = [s.lower() for s in sentences]


def tokenize(s):
    return re.findall(r"\b\w+\b", s)

# building a vocab
tokens = []
for s in sentences: tokens += tokenize(s)
# These are special tokens used during training
special = ["<pad>","<bos>","<eos>","<unk>"]
vocab = special + sorted(set(tokens))
# vocabulary lookup dictionaries
stoi = {w:i for i,w in enumerate(vocab)}

itos = {i:w for w,i in stoi.items()}


# building (X,Y) pairs with window length T
T = 3  # context size (last 3 words -> predict next)
X_list, Y_list = [], []
for s in sentences:
    words = ["<bos>"] + tokenize(s) + ["<eos>"]
    if len(words) < T+1: continue
    for i in range(len(words)-T):
        xw = words[i:i+T]
        yw = words[i+T]
        X_list.append([stoi.get(w, stoi["<unk>"]) for w in xw])
        Y_list.append(stoi.get(yw, stoi["<unk>"]))

X = torch.tensor(X_list, dtype=torch.long)
Y = torch.tensor(Y_list, dtype=torch.long)

# ---------- simple word-level RNN ----------
class WordRNN(nn.Module):
    def __init__(self, vocab_size, embed_dim=32, hidden=64):
        super().__init__()
        self.emb = nn.Embedding(vocab_size, embed_dim)
        self.rnn = nn.RNN(embed_dim, hidden, batch_first=True)
        # self.rnn = nn.LSTM(embed_dim, hidden, batch_first=True) to fix vanishing gradients and let networks remember information over longer spans (long‑term dependencies).
        # self.rnn = nn.GRU(embed_dim, hidden, batch_first=True)  GRU (Gated Recurrent Unit) simplifies LSTM using update and reset gates (fewer parts, no separate cell state).


        self.fc  = nn.Linear(hidden, vocab_size)

    def forward(self, x):
        e = self.emb(x)            # [B, T, E]
        out, _ = self.rnn(e)       # out: [B, T, H]
        last = out[:, -1, :]       # [B, H]
        logits = self.fc(last)     # [B, V]
        return logits

model = WordRNN(vocab_size=len(vocab))
opt = optim.Adam(model.parameters(), lr=0.02)
loss_fn = nn.CrossEntropyLoss()

# mini training loop
batch_size = 32
for epoch in range(20):
    perm = torch.randperm(len(X))
    for i in range(0, len(X), batch_size):
        idx = perm[i:i+batch_size]
        xb, yb = X[idx], Y[idx]
        logits = model(xb)
        loss = loss_fn(logits, yb)
        opt.zero_grad(); loss.backward(); opt.step()
    if (epoch+1) % 5 == 0:
        print(f"epoch {epoch+1}, loss={loss.item():.4f}")

# predict next words given a prompt
def predict(prompt, max_new_tokens=8):
    model.eval()
    words = ["<bos>"] + tokenize(prompt)
    # ensure length T
    if len(words) < T:
        words = ["<bos>"]*(T-len(words)) + words
    else:
        words = words[-T:]
    ctx = torch.tensor([[stoi.get(w, stoi["<unk>"]) for w in words]], dtype=torch.long)
    generated = []
    with torch.no_grad():
        for _ in range(max_new_tokens):
            logits = model(ctx)
            next_id = int(torch.argmax(logits, dim=-1))
            next_word = itos[next_id]
            if next_word == "<eos>":
                break
            generated.append(next_word)
            # slide window
            ctx = torch.cat([ctx[:,1:], torch.tensor([[next_id]], dtype=torch.long)], dim=1)
    return (prompt + " " + " ".join(generated)).strip()

print("Prompt: i like  →", predict("i like", 8))
print("Prompt: deep learning  →", predict("deep learning", 8))
print("Prompt: transformers  →", predict("transformers", 8))


#  Whats happening ?
# It is learning patterns of word sequences from the training data you provided.
# During training, for each sliding window of 3 words (T=3), the model tries to predict the next word.

# Example:
# Input: [ "i", "like", "machine" ] → Target: "learning"
# Input: [ "deep", "learning", "is" ] → Target: "powerful"
# “i like → i like machine learning”
# from our corpus, the contexts after “i like …” are: machine, pizza, coding. The model had learned a distribution and your greedy argmax happened to pick machine first, then it kept going.