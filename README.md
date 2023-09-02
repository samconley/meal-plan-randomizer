## Meal Plan Randomizer
Selects meals at random from a pre-defined list of recipes you provide and sends them to your phone.

## Running the project for yourself
1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo) project.

2. Set the following environment secrets in Github:

- `FROM_EMAIL` -- The email account you want to use as the sender. This may vary depending on your email host. For my purposes, using Gmail, I had to configure an [app password](https://security.google.com/settings/security/apppasswords) to allow the application to send from my Gmail account.

- `EMAIL_PASSWORD` -- The password for the email account above.

- `SMTP_HOST` -- Again, this varies based on your provider. For Gmail, use `smtp.gmail.com`

- `SMTP_HOST_PORT` -- Also provider specific, but typically the port is `465`

- `TO_LIST` -- A comma-separated list of phone numbers to send the text messages to. The general format is `5551234567@xyz.com` where the `xyz.com` suffix varies based on service provider. A list of the more common providers' suffixes can be found [here](https://avtech.com/articles/138/list-of-email-to-sms-addresses/).

3. Optionally, you may edit the environment varibles to suit your needs:

- `LESS_RECENT_THAN_DAYS` -- (Integer) The minimum number of days that must elapse before a meal is eligible to be chosen again.

- `NUM_MEALS_TO_SEND` -- (Integer) The number of meals from your list to be sent each week.

- `SOURCE_FILE_DIR` -- Relative directory where the meals.json file is located.

4. Finally, if you don't like any of my personal meals, edit the source meal list located in etc/meals.json.
Every object within the "meals" array is either a string or an array of strings. If not specified, the "lastUsed"
timestamp will update any time the meal is selected.