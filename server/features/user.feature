Feature: User CRUD operations

Scenario: Adding new user
    Given a user with id 3444 does not exist
    When I add a user with id 3444
    Then the user with id 3444 has to exist

Scenario: Retrieving user by ID
  Given a user with id 3444 exists
  When I get a user with id 3444
  Then the user with id 3444 has to be returned
