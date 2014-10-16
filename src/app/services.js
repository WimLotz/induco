(function () {
    var inducoApi = function ($http) {
        return {
            saveUser: function (user) {
                return $http({method: 'POST', data: user, url: 'http://localhost:4567/saveUser'});
            },
            login: function (user) {
                return $http({method: 'POST', data: user, url: 'http://localhost:4567/login'});
            },
            saveCompanyProfile: function (company) {
                return $http({method: 'POST', data: company, url: 'http://localhost:4567/saveCompanyProfile'});
            },
            fetchCompanyProfiles:function(){
                return $http({method: 'GET', url: 'http://localhost:4567/fetchCompanyProfiles'});
            },
            fetchPersonProfiles:function(){
                return $http({method: 'GET', url: 'http://localhost:4567/fetchPersonProfiles'});
            },
            signOut:function(){
                return $http({method: 'POST', url: 'http://localhost:4567/googleDisconnect'});
            }
        }
    };

    var tagsStorage = function(){
        var tags = [];

        var addTag = function(tag) {
            tags.push(tag);
        }

        var getTags = function(){
            return tags;
        }

        var removeTag = function(index) {
            tags.splice(index, 1);
        }

        return {
            addTag: addTag,
            getTags: getTags,
            removeTag: removeTag
        };
    };

    inducoApi.$inject = ['$http'];
    tagsStorage.$inject = [];

    angular.module("services", [])
        .factory('inducoApi', inducoApi)
        .factory('tagsStorage', tagsStorage);
})();