from lxml import etree ,objectify
import os
from ConfigController import main
import json

file_controller=main()

class XmlParsing():
    
    def __init__(self):
        pass


    def CreateXmlFile(self,file_name,params):
        
        with open(file_controller.FullPath('template')) as file:
            xml_template = file.read()
            file.close()
                
        parameters=params
        root = objectify.fromstring(xml_template)
        root.set("name",file_name)
        counter=0
        
        ## Set Exist Parameters 
        for param in parameters:
                for xml_param in root.settings.param:
                    if(param['name']==xml_param.get("name")):
                        xml_param.set('value',param['value'])
                        del parameters[counter]
                counter+=1
        ## Add New Parameters
        
        for param in parameters:
            root.settings.append(objectify.Element("param",name=param['name'],value=param['value']))  

        return etree.tostring(root)
                
                        
                        
                        
        
        # for i in root.settings.param:
        #     i.set("name","test")

        # for i in root.settings.param:
            
        #     print(i.get("name") + " : "+i.get("value"))
        

