# This file is a template, and might need editing before it works on your project.
from instructure/golang 
COPY goapp.go /home/
CMD ["go","run","/home/goapp.go"]
