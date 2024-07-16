# SendGrid Email Sender

This Go script sends SendGrid dynamic template emails. It reads recipient data from a CSV file and utilizes environment variables for secure configuration of sensitive information.

### Features:
- **SendGrid Integration**: Utilizes SendGrid API to send dynamic template emails.
- **CSV Data Handling**: Reads recipient email addresses and personalized data from a CSV file.
- **Environment Variables**: Stores sensitive information such as API keys and email content in environment variables.

### Usage:
1. **Setup Environment**:
   - Create a `.env` file in the root directory with the following variables:
     ```
     SENDGRID_API_KEY=your_sendgrid_api_key
     TEMPLATE_ID=your_template_id
     FROM_EMAIL=your_email@example.com
     CSV_FILE=emails.csv
     ```

2. **CSV File Format**:
   - Create a CSV file (`emails.csv`) with headers (`email`, `name`, etc.) containing recipient information.

3. **Run the Script**:
   - Execute the script using:
     ```bash
     go run main.go
     ```

### Example `.env` File:
```dotenv
SENDGRID_API_KEY=your_sendgrid_api_key
TEMPLATE_ID=your_template_id
FROM_EMAIL=your_email@example.com
CSV_FILE=emails.csv
