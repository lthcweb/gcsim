import charPipelineData from "./char_data.generated.json";
import weaponPipelineData from "./weapon_data.generated.json";

export interface CharDataMap {
  [key: string]: {
    rarity: number;
    element: string;
    weapon_class: string;
  };
}

export function protoEleToDisplayString(ele: string): string {
  switch (ele) {
    case "Electric":
      return "electro";
    case "Fire":
      return "pyro";
    case "Ice":
      return "cryo";
    case "Water":
      return "hydro";
    case "Grass":
      return "dendro";
    case "Wind":
      return "anemo";
    case "Rock":
      return "geo";
    default:
      return "unknown";
  }
}

export function protoWeapTypeToDisplayString(w: string): string {
  switch (w) {
    case "WEAPON_SWORD_ONE_HAND":
      return "sword";
    case "WEAPON_CLAYMORE":
      return "claymore";
    case "WEAPON_POLE":
      return "polearm";
    case "WEAPON_BOW":
      return "bow";
    case "WEAPON_CATALYST":
      return "catalyst";
    default:
      return "unknown";
  }
}

export const valid_weapons: string[] = Object.keys(weaponPipelineData.data)

let charData: CharDataMap = {};

for (const [k, v] of Object.entries(charPipelineData.data)) {
  charData[k] = {
    rarity: v.rarity === "QUALITY_ORANGE" ? 5 : 4,
    element: protoEleToDisplayString(v.element),
    weapon_class: protoWeapTypeToDisplayString(v.weapon_class),
  };
}

export const CharMap: CharDataMap = charData;

import ArtifactMainStatsData from "./artifact_main_gen.json";

export { ArtifactMainStatsData };
