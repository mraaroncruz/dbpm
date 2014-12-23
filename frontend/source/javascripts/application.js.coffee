angular.module("thepickmachine", [])

angular.module("thepickmachine")
  .controller("ApplicationController", ->
  )
  .controller("PicksController", ($scope, $http) ->
    cachedPicks = []
    baseURL = "http://dbpm.aaroncruz.com"

    $http.get(baseURL).then (res) ->
      $scope.picks = res.data
      cachedPicks = angular.copy(res.data)

    $scope.$watch 'search.term', (curr, prev) ->
      return if curr == prev
      if curr == ""
        $scope.picks = cachedPicks
      else
        $http.get("#{baseURL}/search", params: { q: curr }).then (res) ->
          $scope.picks = res.data

  )
  .directive("logo", ->
    scope:
      pick: "=for"
    replace: true
    template: """
      <span>{{logoSrc}}</span>
    """
    link: (scope, element, attrs) ->
      getLogo = ->
        show = scope.pick.show_name
        letters = (word[0] for word in show.split(" "))
        letters.join('')

      scope.logoSrc = getLogo()
  )
