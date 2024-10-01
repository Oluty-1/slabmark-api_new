package templates

var AccountCreatedTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Forgot Password</title>
    <style>
        body {
            font-family: Helvetica, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 400px;
            margin: 0 auto;
            padding: 20px;
            background-color: #fff;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
        }
        .logo {
            text-align: center;
            padding: 10px 0;
        }
        .logo img {
            max-width: 100px;
            height: auto;
        }
        h1 {
            color: #000;
            font-size: 24px;
            text-align: center;
        }
        .message {
            margin: 20px 0;
            color: #333;
            font-size: 16px;
            line-height: 1.5;
        }
        .button {
            text-align: center;
            margin-top: 20px;
        }
        a.button-link {
            background-color: #007BFF;
            color: #fff;
            padding: 12px 24px;
            text-decoration: none;
            border-radius: 5px;
            display: inline-block;
            transition: background-color 0.3s;
        }
        a.button-link:hover {
            background-color: #0056b3;
        }
        .footer {
            text-align: center;
            margin-top: 20px;
            color: #888;
            font-size: 12px;
        }
    </style>
</head>
<body>
<div class="container">
    <div class="logo">
        <h1 style="color: #007BFF">Slabmark Nig Limited</h1>
    </div>
    <h1>Welcome %s</h1>
    <div class="message">
        <p>We are excited to welcome you as a new user of our integrated tool! Vera is designed to streamline your workflow, enhance productivity, and provide you with an efficient way to manage data, and communication. </p>
    </div>
    <div class="message">
        <p>Department: %s</p>
    </div>
    <div class="message">
        <p>Your password is: <b> %s </b></p>
    </div>
    <div class="message">
        <p>Please login to VERA and change your password</p>
    </div>
    <div class="message">
        <p>If you have any questions or need assistance, please contact our support team at <a href="tejiriaustin123@gmail.com">[Support Email Address]</a>.</p>
    </div>
    <div class="footer">
        &copy; 2023 Slabmark Nig. Limited | Privacy Policy | Unsubscribe
    </div>
</div>
</body>
</html>`
