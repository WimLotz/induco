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

    var tags = function () {
        return {
            restrict: 'E',
            templateUrl: "app/templates/tags.html",
            scope: {
                tagCollectionName: '=?',
                tagCollection: '=?'
            },
            controller: 'TagsController'
        };
    };

    angular.module("directives",[])
        .directive('navigationBar', navigationBar)
        .directive('tags', tags);
})();
