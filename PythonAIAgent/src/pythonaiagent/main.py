from smolagents import CodeAgent
from smolagents.models import OpenAIModel
from dotenv import load_dotenv


def main():
    load_dotenv()
    model = OpenAIModel(model_id="gpt-4.1-nano")  # You can also use "gpt-3.5-turbo"
    agent = CodeAgent(tools=[], model=model)

    result = agent.run("Write a Python function to reverse a string.")
    print("Agent Output:\n", result)

if __name__ == "__main__":
    main()
