import { type Node } from "../../../types/data";
import { create } from "zustand";

type UseNodesState = {
  items: Node[];
  selectedItem: Node | null;
}

type UseNodesActions = {
  setItems: (nodes: Node[]) => void;
  setSelectedItem: (node: Node | null) => void;
}

export const useNodes = create<UseNodesState & UseNodesActions>((set) => ({
  items: [],
  setItems: (nodes) => set({ items: nodes, selectedItem: nodes[0] }),

  selectedItem: null,
  setSelectedItem: (node) => set({ selectedItem: node }),
}))
