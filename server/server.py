from pydantic import BaseModel
from fastapi import FastAPI
from pathlib import Path
import uvicorn

class Message(BaseModel):
    msg: str


def write_to_txt(text):
    Path('../files/').mkdir(parents=True, exist_ok=True)
    with open('../files/test.txt', 'a+') as f:
        f.write(text)
        f.write('\n')


app = FastAPI()

@app.post('/')
async def print_msg(msg: Message):
    write_to_txt(msg.msg)
    return {'msg': 'OK'}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)
