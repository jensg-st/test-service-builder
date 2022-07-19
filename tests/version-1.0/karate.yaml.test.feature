
Feature: Text Function

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def testmeSecret = karate.properties('testmeSecret')
* def jens = karate.properties('jens')


Scenario: get request

	Given url 'http://' + karate.properties('testURL')

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"commands": [
		{
			"command": "ls -la",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 200
		And match $ ==
	"""
	{
	"testme": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""
	