// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateNym } from "./types/verimcosmos/tx";
import { MsgUpdateNym } from "./types/verimcosmos/tx";
import { MsgDeleteNym } from "./types/verimcosmos/tx";


const types = [
  ["/verimid.verimcosmos.verimcosmos.MsgCreateNym", MsgCreateNym],
  ["/verimid.verimcosmos.verimcosmos.MsgUpdateNym", MsgUpdateNym],
  ["/verimid.verimcosmos.verimcosmos.MsgDeleteNym", MsgDeleteNym],
  
];

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgCreateNym: (data: MsgCreateNym): EncodeObject => ({ typeUrl: "/verimid.verimcosmos.verimcosmos.MsgCreateNym", value: data }),
    msgUpdateNym: (data: MsgUpdateNym): EncodeObject => ({ typeUrl: "/verimid.verimcosmos.verimcosmos.MsgUpdateNym", value: data }),
    msgDeleteNym: (data: MsgDeleteNym): EncodeObject => ({ typeUrl: "/verimid.verimcosmos.verimcosmos.MsgDeleteNym", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};