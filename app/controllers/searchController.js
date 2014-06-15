(function () {
    var SearchController = function ($scope) {
        $scope.message = 'last crap message';
    };
    SearchController.$inject = ['$scope'];
    angular.module('induco').controller('SearchController', SearchController);
}());
