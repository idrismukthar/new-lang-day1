import random
import string

def generate_password():
    print("--- 🛡️ SECURE PASSWORD GENERATOR ---")
    
    # Ask the user for their preferences
    length = int(input("Enter desired password length (e.g., 12, 16): "))
    site_name = input("What website is this for? (e.g., Instagram, Binance): ")

    # Define the characters we want to use
    chars = string.ascii_letters + string.digits + "!@#$%^&*"
    
    # The Magic: Randomly pick characters from the list 'length' times
    password = "".join(random.choice(chars) for i in range(length))
    
    print(f"\n✅ Generated Password: {password}")

    # Sensible part: Save it to a file so it's actually useful
    with open("my_vault.txt", "a") as vault:
        vault.write(f"Site: {site_name} | Password: {password}\n")
    
    print("--- Password safely saved to 'my_vault.txt' ---")

# Run the function
generate_password()