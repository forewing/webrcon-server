FROM python:3
ENV PYTHONUNBUFFERED 1
RUN mkdir /srv/django
WORKDIR /srv/django
COPY requirements.txt /srv/django/
RUN python3 -m pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt
