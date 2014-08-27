(function () {

    var inducoApp = angular.module('induco', ['ngRoute']);

    inducoApp.config(function ($routeProvider, $httpProvider) {

        $httpProvider.defaults.useXDomain = true;
        delete $httpProvider.defaults.headers.common['X-Requested-With'];

        $routeProvider
            .when('/',
            {
                controller: 'HomeController',
                templateUrl: 'app/templates/home.html'
            })
            .when('/profile', {
                controller: 'ProfileController',
                templateUrl: 'app/templates/profile.html'
            })
            .when('/search', {
                controller: 'SearchController',
                templateUrl: 'app/templates/search.html'
            })
            .when('/dashboard', {
                controller: 'DashboardController',
                templateUrl: 'app/templates/dashboard.html'
            })
            .otherwise({ redirectTo: '/'});
    });
}());
