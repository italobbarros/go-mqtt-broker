# Cria um cliente MQTT
import asyncio
from aiomqtt import Client
import sys,os
import datetime
import random
import logging

logger = logging.getLogger("Client")
logger.setLevel(logging.DEBUG)
handler = logging.StreamHandler()
formatter = logging.Formatter('%(asctime)s - [%(levelname)s] %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

async def publish_messages(client_id, topic,delay):
    await asyncio.sleep(delay)
    try:
        async with Client("localhost", 1883,client_id=client_id,timeout=60) as client: #broker.hivemq.com
            for _ in range(1): # quantas vezes repetira
                for i in range(1, 101):  # Publica 1000 mensagens em topicos diferentes
                    message = f"Mensagem {i} do {client_id}"
                    topic = f"mqtt/{client_id}/topic{i}"
                    await client.publish(topic, message,qos=1)
                    logger.debug(f"Cliente: {client_id} | topic: {topic} publicou: {message}")
                #await asyncio.sleep(1)
    except Exception as e:
        print(f"Erro de conexão para o Cliente_{client_id}: {e}")

async def main():
    start_time = datetime.datetime.now()
    tasks = []
    # Cria x clientes assíncronos
    for i in range(1, 501):
        clientId = f"clientId{i}"
        tasks.append(asyncio.create_task(publish_messages(clientId, None,0)))
    
    await asyncio.gather(*tasks)
    duration = datetime.datetime.now() - start_time
    logger.debug(f"Tempo total de execução: {duration}")

if __name__ == "__main__":
    # Change to the "Selector" event loop if platform is Windows
    if sys.platform.lower() == "win32" or os.name.lower() == "nt":
        from asyncio import set_event_loop_policy, WindowsSelectorEventLoopPolicy
        set_event_loop_policy(WindowsSelectorEventLoopPolicy())
    # Run your async application as usual
    asyncio.run(main())