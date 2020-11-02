from lxml import etree ,objectify
import os

import json


class main():
    def __init__(self):
        pass


    def FilePath(self):
        __location__ = os.path.realpath(os.path.join(os.getcwd(), os.path.dirname(__file__)))
        file =open(os.path.join(__location__, 'config.json'))
        path =json.load(file)['location']
        return path
    def FullPath(self,filename):
        path=(self.FilePath())+"/"+filename+".xml"
        return (path)

    def ExistOrNo(self,filename):
        path=self.FilePath()
        is_exist=os.path.isfile(path+"/"+filename+".xml")
        return is_exist

