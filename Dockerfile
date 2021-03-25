FROM scratch
MAINTAINER <princegosavi12@gmail.com>

EXPOSE 9090

COPY userMocker /app/userMocker
COPY clientBin /app/clientBin

CMD ["/app/userMocker"]