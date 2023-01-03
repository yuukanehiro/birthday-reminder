module birthday-reminder

go 1.17

replace birthday-reminder/routes => ./routes
replace birthday-reminder/controllers => ./controllers
replace birthday-reminder/packages/Domain/Domain => ./packages/Domain/Domain
replace birthday-reminder/packages/Infrastructure/Repositories/Rdb => ./packages/Infrastructure/Repositories/Rdb

