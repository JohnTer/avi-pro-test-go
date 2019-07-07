import unittest
import re
import requests
import json
import uuid

URL_GET = "http://localhost:8080/api/retrieve/?id=%s"
URL_POST = "http://localhost:8080/api/generate/"


class TestAPI(unittest.TestCase):

    def sendGet(self, url, param):
        url = url % param
        answer = requests.get(url)
        return answer.json()


    def sendPost(self, url, data):
        headers = {'Content-type': 'application/json',
                   'Accept': 'text/plain',
                   'Content-Encoding': 'utf-8'}
        answer = requests.post(url, data=json.dumps(data), headers=headers)
        return answer.json()

    def getRndID(self, url, data):
        return self.sendPost(url, data)["data"]


    def getRndVal(self, url, rnd_id):
        return self.sendGet(url, rnd_id)["data"]


    def getErrFromPost(self, url, data):
        return self.sendPost(url, data)["err"]


    
    def testInvalidLen(self):
        js = """{
  "randType": "str",
  "strLen": "5abc"
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)

    
    def testNoLen(self):
        js = """{
  "randType": "str"
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)

    
    def testNoIntRange(self):
        js = """{
  "randType": "num"
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)
    
    def testInvalidRange(self):
        js = """{
  "randType": "num",
  "intRange": "abc"
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)
    
    def testNoRandType(self):
        js = """{
  "strLen": 52
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)
    
    def testUknownRandType(self):
        js = """{
  "randType": "numabcd"
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)
    
    def testNoExtAlph(self):
        js = """{
  "randType": "extstr",
  "strLen": 5
	}"""
        err = self.getErrFromPost(URL_POST, json.loads(js))
        self.assertEqual(err, 1)


    
    def testRandStr(self):
        js = """{
  "randType": "str",
  "strLen": 5
	}"""
        rand_id = self.getRndID(URL_POST, json.loads(js))
        rand_val = self.getRndVal(URL_GET, rand_id)
        pattern = r'[a-z]{5}'
        self.assertIsNotNone(re.match(pattern, rand_val))
    
    def testRandNum(self):
        js = """{
  "randType": "num",
  "intRange": [1,5]
	}"""
        rand_id = self.getRndID(URL_POST, json.loads(js))
        rand_val = self.getRndVal(URL_GET, rand_id)
        pattern = r'[-+]?\d+'
        self.assertIsNotNone(re.match(pattern, rand_val)) 
    
    def testRandStrNum(self):
        js = """{
  "randType": "strnum",
  "strLen": 5
	}"""
        rand_id = self.getRndID(URL_POST, json.loads(js))
        rand_val = self.getRndVal(URL_GET, rand_id)
        pattern = r'\w{5}'
        self.assertIsNotNone(re.match(pattern, rand_val)) 

    
    def testRandUUID(self):
        js = """{
  "randType": "uuid"
	}"""
        rand_id = self.getRndID(URL_POST, json.loads(js))
        rand_val = self.getRndVal(URL_GET, rand_id)
        uuid.UUID(rand_val)
        
    def testRandExtAlph(self):
        js = """{
  "randType": "extstr",
  "extAlphabet": "abcdef123",
  "strLen": 5
	}"""
        rand_id = self.getRndID(URL_POST, json.loads(js))
        rand_val = self.getRndVal(URL_GET, rand_id)
        pattern = r'[1-3a-f]{5}'
        self.assertIsNotNone(re.match(pattern, rand_val)) 

if __name__ == '__main__':
    unittest.main()
