export type Network = {
  id: string;
  name: string;
  image: string;
  health_status: string;
  progress_status: string;
}

export type Model = {
  id: string;
  name: string;
  image: string;
  memory: number;
}

export type NodeOnchainData = {
  address: string;
  id: string;
  processing_tasks: number;
}

export type Device = {
  id: string;
  name: string;
  os: string;
  processor: string;
  ram: number; // GB
  gpu: string;
  gpu_cores: number;
}

export type Node = {
  id: string;
  name: string;
  image: string;
  earnings: number;
  status: string;
  network: Network;
  model: Model;
  onchain_data: NodeOnchainData;
  device: Device;
}
