from flask import Flask ,current_app, Response,request
from ConfigController import main
from XmlParser import XmlParsing
from django.http import HttpResponse
app = Flask(__name__)

file_controller = main()
xml_config_controller=XmlParsing()




@app.route('/profile',methods=['GET'])
def GetXmlData():
    try:
        name = request.args.get('name')
        if(file_controller.ExistOrNo(name)):
            xml_file = open(file_controller.FullPath(name))
            
            return Response(xml_file,status=200,mimetype='text/xml')
        else:
            return Response("not exitst",status=404)
    except:
        return Response("not found",status=404)




@app.route('/profile',methods=['POST'])
def CreateConfigFile():
    try:
        json_file=request.json
        parameter_data = json_file['params']
        file_name =json_file['profile_name']
        if(file_name!="template"):
            XmlCreated=xml_config_controller.CreateXmlFile(file_name,parameter_data)
            
            with open(file_controller.FullPath(json_file['profile_name']), "wb") as file: 
                file.write(XmlCreated)  
            return Response("ok",status=200)
        else:
            return Response("not allowed . send other profile names",status=405)
    except:
        return Response("not found",status=200)

# @app.route('/profile',methods=['DELETE'])
# def Delete():
#     try:
#         json_file=request.json
#         parameter_data = json_file['params']
#         file_name =json_file['profile_name']
#         XmlCreated=xml_config_controller.CreateXmlFile(file_name,parameter_data)
#         print(XmlCreated)
#         with open(file_controller.FullPath(json_file['profile_name']), "wb") as file: 
#             file.write(XmlCreated)  
#         return Response("ok",status=200)
#     except:
#         return Response("not found",status=200)
if __name__ == '__main__':
    app.run()
