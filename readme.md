# cloudvis
Простой консольный клиент для сервиса [VisionBot](https://visionbot.ru).  
Использование: `cloudvis [опции] файл`  
`Файл` - имя файла в формате png или jpg.  

Поддерживаемые опции:  
`-l language`  
Указывает, какой язык использовать, язык задаётся в виде двухбуквенного кода (например `ru` или `en`). По умолчанию `en`.  
Список поддерживаемых языков см. на сайте сервиса.  
`-stdin`  
Читать изображение с stdin вместо файла.  
`-t type`  
Указывает тип данных для распознавания. Может быть `text`, `image` или `all`. По умолчанию `all`.  
`-tr`  
Перевести результат распознавания. Если флаг установлен, результат будет переведён, иначе вернётся исходный текст. По умолчанию перевод отключён.  
`-qr`  
Распознать QR/BAR коды. Если флаг установлен, то сервис попытается распознать QR/BAR код с изображений. Если флаг не установлен, будет использоваться только стандартное распознавание. По умолчанию распознавание QR/BAR отключено.
## Сборка
Требования:
- Среда разработки [golang](https://golang.org)  

Сборка тестировалась на Go 1.17.6.

Далее выполните такие команды:  
`git clone https://github.com/alekssamos/cloudvis.git`  
`cd cloudvis`  
`go build cloudvis.go`  

Распознавание осуществляется в облаке при использовании [этого](https://visionbot.ru) сервиса.  
Данный консольный клиент не осуществляет какую-либо обработку загружаемых изображений, поэтому обо всех ошибках, связанных с некорректным распознаванием, следует сообщать напрямую [разработчику](mailto:aleks-samos@yandex.ru) сервиса.  
Обратите внимание, что хоть распознавание текста работает достаточно хорошо, не стоит полностью доверять результату в вопросе распознавания предметов.  

В своей работе приложение взаимодействует со следующей конечной точкой API:  
https://visionbot.ru/apiv2

# English
Simple console client for [VisionBot](https://visionbot.ru) service.  
Usage: `cloudvis [opts] file`  
 `file` is a image file name. png and jpg are officially supported.  
 Opts can be one of:  
  `-l language`  
    	Set language name to use. Language is formed as two letter code, like ru or en (default "en")  
  `-stdin`  
    	Get image data from stdin, filename will be ignored  
  `-t type`  
    	Data type to recognize. Can be 'text', 'image' or 'all' (default "all")  
  `-tr`  
    Translate the recognition result. If the flag is set, the result will be translated, otherwise the original text will be returned. Translation is disabled by default  
  `-qr`  
    Recognize QR/BAR codes. If the flag is set, the service will try to recognize the QR/BAR code from the images. If the flag is not set, only standard recognition will be used. QR/BAR recognition is disabled by default  
## Build
Requirements:
- Build environment [golang](https://golang.org)  

Tested on Go 1.17.6.

Next, run the following commands:  
`git clone https://github.com/alekssamos/cloudvis.git`  
`cd cloudvis`  
`go build cloudvis.go`  

All recognition computations are made in the cloud by [this](https://visionbot.ru) service.  
This console client does not perform any processing of uploaded images, so all errors related to incorrect recognition should be reported directly to [the developer](mailto:aleks-samos@yandex.ru) of the service.  
Please note that although text recognition works quite well, you should not completely trust the result in the matter of recognizing objects.  

In its work, the application interacts with the following API endpoint:  
https://visionbot.ru/apiv2