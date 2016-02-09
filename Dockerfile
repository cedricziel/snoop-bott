FROM scratch
ADD snoop-bott /
ENV TELEGRAM_TOKEN ""

CMD ["/snoop-bott]
