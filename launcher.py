import webbrowser # This is a built-in library to control your browser

# A "Dictionary" to store your favorite links
links = {
    "study": [
        "https://github.com/idrismukthar",
        "https://www.w3schools.com/python/",
        "https://stackoverflow.com"
    ],
    "chill": [
        "https://www.youtube.com",
        "https://www.netflix.com",
        "https://open.spotify.com"
    ],
    "news": [
        "https://www.bbc.com",
        "https://www.aljazeera.com"
    ]
}

print("--- 🚀 IDRIS'S GLOBAL LAUNCHER ---")
choice = input("What mode are we in today? (study / chill / news): ").lower()

if choice in links:
    print(f"Opening your {choice} workspace...")
    for url in links[choice]:
        webbrowser.open(url)
else:
    print("Omo, I don't know that mode. Try 'study' or 'chill'.")

print("\nLauncher task complete. Happy coding!")