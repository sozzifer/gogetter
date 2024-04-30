module gogetter

go 1.22.2

require (
    gogetter/stores/stub
    gogetter/models/models
)
replace gogetter/stores/stub => ./stores/stub
