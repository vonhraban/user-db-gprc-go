Feature: User CRUD operations

Scenario: Adding new user
  Given user with id 3444 does not exist
  When I add a user with id 3444
  Then the user with id 3444 has to exist
