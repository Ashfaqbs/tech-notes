import re
import math
import torch
import torch.nn as nn
import torch.optim as optim

torch.manual_seed(0)

# ---------- tiny training corpus (same idea as your RNN script) ----------
sentences = [
    "i like machine learning",
    "i like pizza",
    "i like coding in python",
    "rnn models process sequences",
    "transformers use attention",
    "deep learning is powerful",   
    "pytorch makes building models easy",
    "we predict the next word",
    "this is a simple dataset",
    "neural networks learn patterns"
]
sentences = [s.lower() for s in sentences]

def tokenize(s):
    # simple word tokenizer: sequences of word chars
    return re.findall(r"\b\w+\b", s)

# ---------- vocab ----------
tokens = []
for s in sentences: tokens += tokenize(s)
special = ["<pad>","<bos>","<eos>","<unk>"]
vocab = special + sorted(set(tokens))
stoi = {w:i for i,w in enumerate(vocab)}
itos = {i:w for w,i in stoi.items()}

# ---------- build (X, Y) supervised pairs with context window T ----------
T = 3  # context length (like before)
X_list, Y_list = [], []
for s in sentences:
    words = ["<bos>"] + tokenize(s) + ["<eos>"]
    if len(words) < T+1: 
        continue
    for i in range(len(words)-T):
        ctx = words[i:i+T]
        nxt = words[i+T]
        X_list.append([stoi.get(w, stoi["<unk>"]) for w in ctx])
        Y_list.append(stoi.get(nxt, stoi["<unk>"]))

X = torch.tensor(X_list, dtype=torch.long)  # [N, T]
Y = torch.tensor(Y_list, dtype=torch.long)  # [N]

# ---------- tiny GPT-ish Transformer block ----------
class TinyTransformer(nn.Module):
    def __init__(self, vocab_size, d_model=64, nhead=4, num_layers=2, max_T=64, dim_ff=128):
        super().__init__()
        self.tok_emb = nn.Embedding(vocab_size, d_model)
        self.pos_emb = nn.Embedding(max_T, d_model)  # learned positional embeddings
        layer = nn.TransformerEncoderLayer(
            d_model=d_model, nhead=nhead, dim_feedforward=dim_ff, batch_first=True
        )
        self.encoder = nn.TransformerEncoder(layer, num_layers=num_layers)
        self.lm_head = nn.Linear(d_model, vocab_size)

    def forward(self, x):
        # x: [B, T] token ids (only past context, so causal masking is optional here)
        B, T = x.shape
        tok = self.tok_emb(x)                               # [B, T, d_model]
        pos = self.pos_emb(torch.arange(T, device=x.device))# [T, d_model]
        h = tok + pos                                        # [B, T, d_model]

        # Build a causal mask so token t can't attend to future tokens (>t)
        # encoder's src_key_padding_mask is for pads, not causality; we use a square subsequent mask.
        # PyTorch's encoder expects [T, T] mask with True/-inf semantics depending on API;
        # the layer uses attn_mask where True/float(-inf) blocks. We'll use float mask.
        attn_mask = torch.full((T, T), float("-inf"), device=x.device)
        attn_mask = torch.triu(attn_mask, diagonal=1)  # upper-triangular is -inf (block future)

        # BUT TransformerEncoder(batch_first=True) still expects attn_mask shape [T, T]
        h = self.encoder(h, mask=attn_mask)

        last = h[:, -1, :]                  # take last position's hidden
        logits = self.lm_head(last)         # [B, vocab]
        return logits

model = TinyTransformer(vocab_size=len(vocab))
opt = optim.Adam(model.parameters(), lr=0.02)
loss_fn = nn.CrossEntropyLoss()

# ---------- training ----------
batch_size = 32
for epoch in range(25):
    perm = torch.randperm(len(X))
    for i in range(0, len(X), batch_size):
        idx = perm[i:i+batch_size]
        xb, yb = X[idx], Y[idx]
        logits = model(xb)
        loss = loss_fn(logits, yb)
        opt.zero_grad(); loss.backward(); opt.step()
    if (epoch+1) % 5 == 0:
        print(f"epoch {epoch+1}, loss={loss.item():.4f}")

# ---------- generation (greedy by default; can toggle temperature/sampling) ----------
def predict(prompt, max_new_tokens=8, temperature=0.0):
    model.eval()
    words = ["<bos>"] + tokenize(prompt)
    # ensure context length T
    if len(words) < T:
        words = ["<bos>"]*(T-len(words)) + words
    else:
        words = words[-T:]
    ctx = torch.tensor([[stoi.get(w, stoi["<unk>"]) for w in words]], dtype=torch.long)

    generated = []
    with torch.no_grad():
        for _ in range(max_new_tokens):
            logits = model(ctx)  # [1, V]
            if temperature and temperature > 0:
                probs = torch.softmax(logits / temperature, dim=-1)
                next_id = int(torch.multinomial(probs, num_samples=1))
            else:
                next_id = int(torch.argmax(logits, dim=-1))
            next_word = itos[next_id]
            if next_word == "<eos>":
                break
            generated.append(next_word)
            # slide window forward
            ctx = torch.cat([ctx[:,1:], torch.tensor([[next_id]], dtype=torch.long)], dim=1)
    return (prompt + " " + " ".join(generated)).strip()

print("Prompt: i like  →", predict("i like", 8))
print("Prompt: deep learning  →", predict("deep learning", 8))
print("Prompt: transformers  →", predict("transformers", 8))
print("Prompt (sampling, T=0.8): i like  →", predict("i like", 8, temperature=0.8))
