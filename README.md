
---

# Expense Tracker CLI

A lightweight and efficient Command-Line Expense Tracker written in Go, designed to help you record and manage daily expenses directly from your terminal.
All your data is stored locally in a JSON file, ensuring simplicity and full control over your personal expense records.

---

## Features

* **Add expenses** with descriptions, amounts, and optional tags
* **List expenses** by tags or by month
* **Summarize** total expenses overall or for a specific month
* **Set and view budget** for any month
* **Export data** to CSV format
* **Delete expenses** by ID

---

## Installation

Follow these steps to install and run the Expense Tracker:

```bash
# 1. Clone the repository
git clone https://github.com/zxCroshka/Expense-Tracker.git

# 2. Navigate to the project directory
cd Expense-Tracker

# 3. Initialize Go module
go mod init github.com/zxCroshka/Expense-Tracker

# 4. Build the project
go build .

# 5. Create a shortcut (alias) for easy use
alias expense-tracker='./Expense-Tracker'
```

---

## Usage

Here are some examples of how to use the **Expense Tracker CLI**:

### 🧾 Add expenses

```bash
expense-tracker add --description "Breakfast" --amount 40
expense-tracker add --description "Lunch" --amount 20 --tags food cheap dish
expense-tracker add --description "Dinner" --amount 10 --tags dish
```

### List expenses

```bash
expense-tracker list
expense-tracker list --tags dish
expense-tracker list --month 10 --tags food
```

### 📊 View summary

```bash
expense-tracker summary
expense-tracker summary --month 10
```

### Manage budget

```bash
expense-tracker budget
expense-tracker config-budget --budget 70
expense-tracker budget --month 10
```

### Export data

```bash
expense-tracker csv
```

This will export your expense data to a `.csv` file.

### Delete an expense

```bash
expense-tracker delete --id 2
```

---

## Example Output

**Listing expenses:**

```
ID   Date        Description  Amount  Categories

1   2025-10-31   Breakfast    40      
2   2025-10-31   Lunch        20      food, cheap, dish
3   2025-10-31   Dinner       10      dish
```

**Summary:**

```
Total expenses this month: 70
Budge for this montht: 0
```

---

## 🧑‍💻 Tech Stack

* **Language:** Go
* **Type:** CLI Application
* **Persistence:** JSON
* **Export Format:** CSV

---

