(function () {
    var DashboardController = function ($scope, $http) {
        $scope.message = "dashboard page"

        $scope.test = function () {
            $http({method: 'GET', url: 'http://localhost:4567/dashboard'})
                .success(function (data) {
                    console.log(data);
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };
    };

    DashboardController.$inject = ['$scope', '$http'];
    angular.module('induco').controller('DashboardController', DashboardController);
}());