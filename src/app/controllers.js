(function() {

    var TagsController = function($scope) {

        $scope.addTag = function() {
            $scope.tagCollection.push({
                name: $scope.tagName
            });
        };

        $scope.remove = function(tagName) {
            for (var i = 0; i < tagName.length; i++) {
                if ($scope.tagCollection[i].name === tagName) {
                    $scope.tagCollection.splice(i, 1);
                    return;
                }
            }
        };
    };

    var DashboardController = function($scope) {
        $scope.message = "dashboard page";
    };

    var HomeController = function($scope, $modal, inducoApi, $location) {
        $scope.user = {};

        $scope.login = function() {
            inducoApi.login($scope.user).success(function() {
                $location.path('profile');
                //todo comes in here even if 500 for no user
            });
        };

        $scope.open = function() {
            $modal.open({
                templateUrl: 'app/modals/create_user.html',
                controller: CreateUserController
            });
        };
    };

    var CreateUserController = function($scope, $modalInstance, inducoApi) {
        $scope.user = {};

        $scope.ok = function() {
            inducoApi.saveUser($scope.user);
            $modalInstance.close();
        };

        $scope.cancel = function() {
            $modalInstance.dismiss('cancel');
        };
    };

    var IndexController = function($scope, $location) {
        $scope.isHomePage = function() {
            return $location.path() === '/';
        };
    };

    var NavigationBarController = function($scope, inducoApi) {
        $scope.signOut = function() {
            inducoApi.signOut();
        };
    };

    var SearchController = function($scope) {
        $scope.message = 'search page';
    };

    var ProfileController = function($scope, inducoApi, tagsStorages, $location) {
        $scope.profiles = [];
        $scope.personProfiles = [];
        $scope.companyProfiles = [];

        inducoApi.fetchUserProfiles().success(function(data) {
            data.forEach(function(profile) {
                if (profile.isCompany) {
                    $scope.companyProfiles.push(profile);
                } else {
                    $scope.personProfiles.push(profile);
                }
                $scope.profiles.push(profile);
            });
        });

        $scope.createPersonProfile = function() {
            $location.path("createPersonalProfile");
        };

        $scope.createCompanyProfile = function() {
            $location.path("createCompanyProfile");
        };
    };

    var CreatePersonProfileController = function($scope, inducoApi) {
        $scope.personProfile = {
            workExpTags: [],
            workExpNeededTags: []
        };
        $scope.showRequiredFields = false;

        $scope.savePersonProfile = function(personProfileForm) {
            if (personProfileForm.$valid) {
                $scope.personProfile.IsCompany = false;
                inducoApi.saveProfile($scope.personProfile);
                $scope.showRequiredFields = false;
            } else {
                $scope.showRequiredFields = true;
            }
        };
    };

    var CreateCompanyProfileController = function($scope, inducoApi) {
        $scope.companyProfile = {};

        $scope.submitCompanyProfileForm = function() {
            $scope.companyProfile.IsCompany = true;
            inducoApi.saveProfile($scope.companyProfile);
        };
    };

    ProfileController.$inject = ['$scope', 'inducoApi', 'tagsStorage', '$location'];
    DashboardController.$inject = ['$scope'];
    HomeController.$inject = ['$scope', '$modal', 'inducoApi', '$location'];
    NavigationBarController.$inject = ['$scope', 'inducoApi'];
    SearchController.$inject = ['$scope'];
    TagsController.$inject = ['$scope'];
    CreateUserController.$inject = ['$scope', '$modalInstance', 'inducoApi'];
    CreatePersonProfileController.$inject = ['$scope', 'inducoApi'];
    CreateCompanyProfileController.$inject = ['$scope', 'inducoApi'];
    IndexController.$inject = ['$scope', '$location'];

    angular.module("controllers", [])
        .controller('DashboardController', DashboardController)
        .controller('HomeController', HomeController)
        .controller('NavigationBarController', NavigationBarController)
        .controller('SearchController', SearchController)
        .controller('TagsController', TagsController)
        .controller('CreateUserController', CreateUserController)
        .controller('CreatePersonProfileController', CreatePersonProfileController)
        .controller('CreateCompanyProfileController', CreateCompanyProfileController)
        .controller('IndexController', IndexController)
        .controller('ProfileController', ProfileController);
})();