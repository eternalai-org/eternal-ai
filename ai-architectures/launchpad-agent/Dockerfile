from python:3.12-alpine
workdir /workspace

copy requirements.txt ./
run --mount=type=cache,target=/root/.cache/pip python -m pip install -r requirements.txt

copy app ./app
copy server.py ./

expose 80
entrypoint ["python", "server.py"]