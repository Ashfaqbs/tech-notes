Maven rough notes ;


---

Maven = build automation + dependency manager.

Core responsibilities:

* Compile source
* Resolve dependencies
* Run tests
* Package artifact (jar/war)
* Install to local repo
* Deploy to remote repo

Convention over configuration:
Standard directory layout:
src/main/java
src/test/java
pom.xml

Maven assumes structure. That’s why minimal config works.

---

2. Maven Lifecycle

---

Lifecycle = predefined sequence of phases.

Main lifecycles:

* clean
* default
* site

Most work happens in default lifecycle.

Default lifecycle phases (important ones):

validate
compile
test
package
verify
install
deploy

Each phase runs all previous phases.

Example:

mvn package
→ runs validate → compile → test → package

mvn install
→ runs validate → compile → test → package → install

mvn deploy
→ runs everything → deploy

Key idea:
A later phase includes earlier phases automatically.

---

3. Clean Lifecycle

---

mvn clean

Deletes target/ directory.

Does not compile.
Does not resolve dependencies beyond what’s needed for clean plugin.

Use when:

* Want fresh build
* Build artifacts corrupted
* Switching branches

---

4. Common Maven Commands

---

mvn compile
→ Compiles main code.

mvn test
→ Compiles + runs tests.

mvn package
→ Builds artifact (jar/war).

mvn install
→ Builds + installs artifact into local repository (~/.m2).

mvn deploy
→ Publishes artifact to remote repository.

mvn clean package
→ Clean + full build.

mvn dependency:tree
→ Show full dependency graph.

---

5. Dependency Resolution

---

Dependencies declared in pom.xml:

<dependency>
   groupId
   artifactId
   version
   scope
</dependency>

Resolution steps:

1. Check local repo (~/.m2/repository)
2. If not found → download from remote (Maven Central)
3. Cache locally

Downloaded once per version.

---

6. Dependency Scopes

---

Scope defines:

* When dependency is available
* Whether it is packaged inside final artifact
* Whether it is visible to tests

Scopes:

1. compile (default)

Available:

* main code
* test code
  Packaged:
* yes

Use when:
Core library required for application runtime.

Examples:
Jackson
Guava
Most business dependencies

This is the default if no scope specified.

---

2. provided

Available:

* compile
* test
  Packaged:
* NO

Use when:
Dependency is provided by runtime/container.

Examples:
Servlet API in Tomcat
Flink core when running on cluster
JEE APIs

Why:
Avoid duplicating runtime libraries inside fat jar.

---

3. runtime

Available:

* runtime
* test
  NOT available:
* compile

Packaged:

* yes

Use when:
Not needed at compile time, but required at runtime.

Example:
JDBC driver sometimes
Logging implementation

Compile code depends only on interface; implementation needed only at runtime.

---

4. test

Available:

* test only
  Not available:
* main code

Packaged:

* no

Use when:
Testing frameworks and test utilities.

Examples:
JUnit
Mockito

---

5. system (rare, discouraged)

Requires manual path.
Not resolved from repository.
Avoid in modern builds.

---

6. import (special, used in dependencyManagement)

Used to import BOM (Bill of Materials).

Example:
Spring Boot BOM
Flink BOM

Manages versions centrally.

---

7. Transitive Dependencies

---

If A depends on B,
and B depends on C,
then A automatically gets C.

Maven resolves full dependency graph.

Use:
mvn dependency:tree

Important for:
Conflict debugging.

---

8. Packaging Types

---

jar (default)
war
pom (parent project)

pom packaging used for:

* Multi-module builds
* Dependency management

---

9. Plugins

---

Maven core does almost nothing.
Plugins do actual work.

Examples:
maven-compiler-plugin
maven-surefire-plugin
maven-shade-plugin

Each plugin binds goals to lifecycle phases.

Example:
compiler plugin → compile phase
surefire → test phase
shade → package phase

---

10. Shade Plugin

---

Used to create fat jar.

Bundles all dependencies inside single jar.

Important for:
Flink jobs
Standalone applications

Without shade:
Only project classes go in jar.
Dependencies must exist externally.

---

11. Local Repository

---

Location:
~/.m2/repository

Stores:
Downloaded dependencies
Installed local artifacts

mvn install copies artifact here.

---

12. Multi-module Concept

---

Parent pom (packaging=pom)
Child modules
Single build across modules

Use when:
Shared libraries
Layered architecture
Large systems

---

13. When to use what

---

During development:
mvn compile
mvn test

Before commit:
mvn clean package

If artifact needed by another local project:
mvn install

CI pipeline:
mvn clean verify
mvn deploy

---
