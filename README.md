````markdown
# q-dep-updater 🧠🔧

A smart CLI tool for **automated, AI-assisted Go module dependency upgrades** — powered by [Amazon Q Developer](https://docs.aws.amazon.com/q/developer/).  
It safely scans, upgrades, tests, and opens a pull request — all without breaking your app.

---

## 🚀 Features

- ✅ Scans your Go project for outdated dependencies  
- 🤖 Uses Amazon Q Developer to analyze code for compatibility issues  
- ⚠️ Predicts breaking changes *before* upgrading  
- 🧪 Runs test suites after upgrades  
- 🔄 Creates a Git feature branch and opens a PR if tests pass  

---

## 📦 Installation

```bash
git clone https://github.com/jesee-kuya/cli-rules
cd q-dep-updater
go build -o q-dep-updater
````

---

## ⚙️ Usage

```bash
./q-dep-updater -path=/absolute/path/to/project -repo=yourusername/projectname
```

### Flags

| Flag       | Description                               | Required |
| ---------- | ----------------------------------------- | -------- |
| `-path`    | Absolute path to your Go project          | ✅ Yes    |
| `-repo`    | GitHub repo (format: user/repo)           | ✅ Yes    |
| `-dryrun`  | Skip making actual changes (just preview) | ❌ No     |
| `-verbose` | Output debug logs                         | ❌ No     |

---

## 🧠 Powered by Amazon Q Developer

This tool uses Amazon Q Developer's CLI to:

* Analyze your `go.mod` and source files
* Predict how a dependency upgrade might break your app
* Recommend code changes to ensure smooth upgrades

---

## ✅ Example

```bash
./q-dep-updater -path=~/projects/url-shortener -repo=myuser/url-shortener
```

**What it does:**

1. Detects outdated modules
2. Analyzes your code with Amazon Q for upgrade safety
3. Runs `go get -u` for safe packages
4. Executes `go test ./...`
5. Creates a GitHub PR like:

   ```
   Title: chore: safe dependency upgrades  
   Body: Auto-upgraded dependencies with passing tests. Verified by Amazon Q Developer.
   ```

---

## 🛠 Requirements

* Go 1.19+
* [Amazon Q Developer CLI](https://docs.aws.amazon.com/q/developer/getting-started/)
* GitHub CLI (`gh` command)
* Git installed and configured

---

## 🧪 Testing Credentials

If you’re testing the app with a sample repo or staging environment:

* **Admin:**
  `username: admin`
  `password: 2025DEVChallenge`

* **User:**
  `username: newuser`
  `password: 2025DEVChallenge`

---

## 📄 License

MIT © 2025 \[jesee-kuya]

---

## 🤝 Contributing

PRs are welcome! For major changes, please open an issue first to discuss what you’d like to change.

---

## 🌍 Connect

📫 linkedIn: [Jesee (Jackson) Kuya](www.linkedin.com/in/jeseekuya)
🔗 Blog post: [Dev.to Challenge Submission](https://dev.to/jeseekuya/zero-downtime-go-module-updates-using-ai-and-cli-automation-11e8)

```
