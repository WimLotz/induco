(function () {

    var navigationBar = function () {
        return {
            restrict: 'E',
            templateUrl: "app/templates/navigation.html",
            scope: {
                activeNavElement: '=?'
            },
            controller: 'NavigationBarController'
        };
    };

    angular.module("directives",[]).directive('navigationBar', navigationBar);
})();
