(function () {
    var HomeController = function ($scope, $http, $window) {
        console.log("headers:" + $http)
        $scope.signIn = function () {

            $http.defaults.useXDomain = true;

            $http({method: 'POST', url: 'http://localhost:8080/googleAuthorise'})
                .success(function (data, status, headers, config) {
                    $window.location.href = data.url;
                    console.log(data)
                })
                .error(function (data, status, headers, config) {
                    console.log('error');
                    console.log(status);
                    console.log(headers);
                    console.log(config);
                });
        };

        $scope.test= function(){
            $http({method: 'GET', url: 'http://localhost:8080/test'})
                .success(function (data, status, headers, config) {
                    console.dir("headers" + headers())
                    console.log(data)
                })
                .error(function (data, status, headers, config) {
                    console.log('error');
                    console.log(status);
                    console.log(headers);
                    console.log(config);
                });
        };
    };
    HomeController.$inject = ['$scope', '$http', '$window'];
    angular.module('induco').controller('HomeController', HomeController);
}());