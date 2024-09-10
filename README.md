# GoExpert-UnitOfWork

This project is a simple example of how to use Unit of Work (UOW) pattern in Go.

For more information on the Unit of Work pattern, see [Martin Fowler's article](https://martinfowler.com/eaaCatalog/unitOfWork.html).

## Run

This project is not a full-fledged application, but a simple example of how to use the Unit of Work pattern.

To run the project, you need to have Docker and Docker Compose installed.

```bash
make up
```

Run the tests in `add_course_test.go`, try changing the category ID to a non-existent one and see what happens. You should see an error message about foreign key constraint violation. You will see that the category was created, but the course was not.

```bash
# Connect to the MySQL container and check the data
docker-compose exec mysql bash
mysql -u root -p courses # password: root
SELECT * FROM categories;
SELECT * FROM courses;
```

On the other hand, if you run the tests with UOW in `add_course_uow_test.go`, if the course cannot be created, then the category will not be created either. This is because the UOW will rollback the transaction, so no changes will be made in the database.

You can check the Unit of Work implementation in `pkg/uow/uow.go`.
