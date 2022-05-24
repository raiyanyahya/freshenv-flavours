from ubuntu:18.04
LABEL maintainer="Raiyan Yahya <raiyanyahyadeveloper@gmail.com>"

RUN apt update -y &&  apt-get update -y
RUN apt-get install --no-install-recommends -y sudo wget git curl zsh wget nano lsof software-properties-common && \
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
