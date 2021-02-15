FROM alpine
ADD works /works
ENTRYPOINT [ "/works" ]
