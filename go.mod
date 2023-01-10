module birthday-reminder

go 1.19


replace birthday-reminder/config => ./config
replace birthday-reminder/routes => ./routes
replace birthday-reminder/controllers => ./controllers

// BirthDay
replace birthday-reminder/packages/Domain/Domain/BirthDay => ./packages/Domain/Domain/BirthDay
replace birthday-reminder/packages/Infrastructure/Repositories/BirthDay => ./packages/Infrastructure/Repositories/BirthDay
replace birthday-reminder/packages/UseCase/BirthDay/Create => ./packages/UseCase/BirthDay/Create
replace birthday-reminder/packages/UseCase/BirthDay/List => ./packages/UseCase/BirthDay/List
replace birthday-reminder/packages/Domain/Application/BirthDay => ./packages/Domain/Application/BirthDay
// Auth
replace "birthday-reminder/packages/Domain/Domain/Auth" => ./packages/Domain/Domain/Auth
replace birthday-reminder/packages/UseCase/Auth/Register => ./packages/UseCase/Auth/Register
replace birthday-reminder/packages/Domain/Application/Auth => ./packages/Domain/Application/Auth
replace birthday-reminder/packages/Infrastructure/Repositories/Auth => ./packages/Infrastructure/Repositories/Auth
// RDB
replace birthday-reminder/packages/Domain/Domain/Rdb => ./packages/Domain/Domain/Rdb
replace birthday-reminder/packages/Infrastructure/Repositories/Rdb => ./packages/Infrastructure/Repositories/Rdb
replace birthday-reminder/packages/MockInfrastructure/Repositories/Rdb => ./packages/MockInfrastructure/Repositories/Rdb
// other
replace birthday-reminder/packages/Domain/Domain/Empty => ./packages/Domain/Domain/Empty
replace birthday-reminder/packages/Domain/Domain/Timestamp => ./packages/Domain/Domain/Timestamp
replace birthday-reminder/packages/Domain/Domain/Date => ./packages/Domain/Domain/Date
replace birthday-reminder/packages/Domain/Domain/Datetime => ./packages/Domain/Domain/Datetime
