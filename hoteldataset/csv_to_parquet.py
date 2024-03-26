#!/usr/bin/env python3

import polars as pl

def convert_csv_to_parquet(csv_file, parquet_file):
    df = pl.read_csv(csv_file)
    df.write_parquet(parquet_file)

if __name__=="__main__":
    print("Converting CSV to Parquet...")
    for i in range(1, 11):
        csv_file = f"dataset_hotel_{i}.csv"
        parquet_file = f"dataset_hotel_{i}.parquet"
        convert_csv_to_parquet(csv_file, parquet_file)