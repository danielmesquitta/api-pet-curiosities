# Daily Pet Curiosity: All About Your Pet!

- Instead of views table, change it to be user_curiosities with viewed (boolean)
- Update business logic when creating curiosities and viewing it:
  - View should be done in the list_daily_curiosities use case
  - After creating new curiosity, it should be linked to an user through the user_curiosities table
- Sending notifications should be done through the user_curiosities table, after the cron job creating curiosities
