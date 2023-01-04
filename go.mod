module birthday-reminder

go 1.19

replace birthday-reminder/packages/Infrastructure/Repositories/Rdb => ./packages/Infrastructure/Repositories/Rdb
replace birthday-reminder/packages/MockInfrastructure/Repositories/Rdb => ./packages/MockInfrastructure/Repositories/Rdb

replace birthday-reminder/packages/Domain/Domain/Rdb => ./packages/Domain/Domain/Rdb

replace birthday-reminder/controllers => ./controllers

replace birthday-reminder/routes => ./routes


