from debian:11
LABEL maintainer="Raiyan Yahya <raiyanyahyadeveloper@gmail.com>"

COPY --from=golang /usr/local/go/ /usr/local/go/  
ENV PATH="/usr/local/go/bin:${PATH}" 

RUN apt update -y && apt-get update -y && apt-get install --no-install-recommends -y sudo git curl zsh wget nano lsof && \
        rm -rf /var/lib/apt/lists/* && \
	apt-get clean && \
        apt-get autoclean && \
        apt-get autoremove
USER root
ENV TERM xterm
WORKDIR /home
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh
RUN chmod -R 755 /usr/local/share/zsh/site-functions

CMD ["zsh"]
