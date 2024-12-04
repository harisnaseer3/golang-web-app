mkdir golang-web-app\db
mkdir golang-web-app\models
mkdir golang-web-app\handlers
mkdir golang-web-app\middleware
mkdir golang-web-app\utils
mkdir golang-web-app\migrations

type nul > golang-web-app\main.go
type nul > golang-web-app\db\connection.go
type nul > golang-web-app\models\user.go
type nul > golang-web-app\models\patient.go
type nul > golang-web-app\handlers\auth.go
type nul > golang-web-app\handlers\patients.go
type nul > golang-web-app\middleware\auth.go
type nul > golang-web-app\utils\jwt.go
type nul > golang-web-app\utils\hash.go
type nul > golang-web-app\migrations\init.sql

echo Project structure created successfully!
