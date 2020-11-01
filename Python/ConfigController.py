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





    # def Reading(self):


       
        # # how to extract element data
        # begin = root.appointment.begin
        # uid = root.appointment.uid

        # # loop over elements and print their tags and text
        # for appt in root.getchildren():
        #     for e in appt.getchildren():
        #         print("%s => %s" % (e.tag, e.text))
        #     print()

        # # how to change an element's text
        # root.appointment.begin = "something else"
        # print(root.appointment.begin)

        # # how to add a new element
        # root.appointment.new_element = "new data"

        # # remove the py:pytype stuff
        # objectify.deannotate(root)
        # etree.cleanup_namespaces(root)
        # obj_xml = etree.tostring(root, pretty_print=True)
        # print(obj_xml)
