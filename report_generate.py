import pandas as pd
import argparse
from ast import literal_eval


def main():
    parser = argparse.ArgumentParser()

    parser.add_argument("-js", "--json", dest="json", default="[{}]", help="Json data")
    parser.add_argument("-ds", "--destination", dest="destination", default="", help="output.xlsx")

    args = parser.parse_args()

    pd.DataFrame.from_dict(literal_eval(args.json)).to_excel(args.destination)
    print("file " + args.destination + " generated")


if __name__ == "__main__":
    main()
