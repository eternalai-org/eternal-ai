from python:3.12-alpine
workdir /workspace

copy app ./app
copy utils ./utils
copy server.py requirements.txt ./

run --mount=type=cache,target=/root/.cache/pip python -m pip install -r requirements.txt
expose 80

entrypoint ["python", "server.py"]