import time

# List of the languages you've touched so far
languages_done = ["C++", "PHP"]
current_lang = "Python"

print("--- 🚀 10 LANGUAGES CHALLENGE: DAY 2 ---")
print(f"Status: You've finished {len(languages_done)} languages.")
print(f"Today's Mission: Master {current_lang} syntax.")

# A cool little loading animation to show Python's speed
print("\nBooting up Python engine...", end="")
for _ in range(5):
    time.sleep(0.3)
    print(".", end="", flush=True)

print("\n\n✅ Python is running perfectly!")

# Let's save your progress to a file (MUCH easier than PHP)
with open("challenge_log.txt", "a") as file:
    file.write(f"Day 2: Successfully installed and ran Python 3.14\n")

print("Progress saved to 'challenge_log.txt'.")