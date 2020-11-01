from flask import Flask ,current_app, Response,request
from ConfigController import main
from django.http import HttpResponse
app = Flask(__name__)

config_controller = main()


@app.route('/profile',methods=['GET'])
def GetXmlData():
    
    name = request.args.get('name')
    if(config_controller.ExistOrNo(name)):
        xml_file = open(config_controller.FullPath(name))
        
        return Response(xml_file,status=200,mimetype='text/xml')
    else:
        return Response("not exitst",status=404)

if __name__ == '__main__':
    app.run()
