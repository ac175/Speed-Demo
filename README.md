
Kirk Miller <kirkpmiller@hotmail.com>
5:53 PM (0 minutes ago)
to me

# ⚡️ Speed-Demo

> **One-file audit-hash micro-service**  
> hashes log lines faster than you can scroll — no database, no external deps.

---

## 1 • Why this repo exists
High-frequency systems (trading, IoT, cloud runtimes) generate millions of events per minute.  
Auditors need **tamper-evident hashes** that keep up without adding latency or infra cost.  
Speed-Demo proves you can hit **≥ 100 000 lines / sec on a laptop core** with nothing but Go.

---

## 2 • Quick start

```bash
# clone & build
git clone https://github.com/YOURORG/Speed-Demo.git
cd Speed-Demo
make build         # static binary appears at ./audit-hash

# hash a single log line
echo '{"event":"login","ts":"2025-05-07T19:15:00Z"}' | ./audit-hash
