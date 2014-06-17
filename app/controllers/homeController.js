(function () {
    var HomeController = function ($scope, $http, $window) {
        $scope.isLoggedIn = false;

        $http({method: 'GET', url: 'http://localhost:8080/isLoggedIn'})
            .success(function (data) {
                console.log(data)
                $scope.isLoggedIn = data.isLoggedIn;
            })
            .error(function (data) {
                    console.log('An error has occurred: ' + data);
            });


        $scope.signIn = function () {
            $http({method: 'POST', url: 'http://localhost:8080/googleAuthorise'})
                .success(function (data) {
                    $window.location.href = data.url;
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };
    };

    HomeController.$inject = ['$scope', '$http', '$window'];
    angular.module('induco').controller('HomeController', HomeController);
}());