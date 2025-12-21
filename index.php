<?php
$filename = 'guestbook.txt';

// 1. Logic to save the message when the form is submitted
if ($_SERVER['REQUEST_METHOD'] == 'POST' && !empty($_POST['message'])) {
    $name = htmlspecialchars($_POST['name']); // Sanitize input
    $message = htmlspecialchars($_POST['message']);
    $entry = "<strong>" . $name . "</strong>: " . $message . "<br>\n";
    
    // Append the message to our text file
    file_put_contents($filename, $entry, FILE_APPEND);
}

// 2. Read existing messages
$entries = file_exists($filename) ? file_get_contents($filename) : "No messages yet!";
?>

<!DOCTYPE html>
<html>
<head>
    <title>PHP Guestbook</title>
</head>
<body>
    <h1>My PHP Guestbook</h1>
    
    <form method="POST">
        <input type="text" name="name" placeholder="Your Name" required><br><br>
        <textarea name="message" placeholder="Write something..." required></textarea><br>
        <button type="submit">Post to Guestbook</button>
    </form>

    <hr>

    <h3>Recent Messages:</h3>
    <div>
        <?php echo $entries; ?>
    </div>
</body>
</html>