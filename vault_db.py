import sqlite3

# 1. Connect to the database (it creates the file if it doesn't exist)
connection = sqlite3.connect("passwords.db")
cursor = connection.cursor()

# 2. Create a Table using SQL syntax
cursor.execute("""
CREATE TABLE IF NOT EXISTS vault (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    site TEXT,
    password TEXT
)
""")

print("--- 🗄️ SQL DATABASE VAULT ---")
site = input("Enter Website: ")
pwd = input("Enter Password: ")

# 3. Insert data using SQL
cursor.execute("INSERT INTO vault (site, password) VALUES (?, ?)", (site, pwd))

# 4. Save and Show
connection.commit()
print("\n✅ Data saved to SQL database!")

print("\n--- Current Vault Entries ---")
cursor.execute("SELECT * FROM vault")
for row in cursor.fetchall():
    print(f"ID: {row[0]} | Site: {row[1]} | Password: {row[2]}")

connection.close()
