// Code generated by plangen.

// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queries

var TpccPlanTests = []QueryPlanTest{
	{
		Query: `
-- cycle 1a
SELECT c_discount, c_last, c_credit, w_tax FROM customer2, warehouse2 WHERE w_id = 1 AND c_w_id = w_id AND c_d_id = 9 AND c_id = 2151`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [customer2.c_discount:7, customer2.c_last:5, customer2.c_credit:6, warehouse2.w_tax:1]\n" +
			" └─ LookupJoin\n" +
			"     ├─ IndexedTableAccess(warehouse2)\n" +
			"     │   ├─ index: [warehouse2.w_id]\n" +
			"     │   ├─ static: [{[1, 1]}]\n" +
			"     │   ├─ colSet: (22-30)\n" +
			"     │   ├─ tableId: 2\n" +
			"     │   └─ Table\n" +
			"     │       ├─ name: warehouse2\n" +
			"     │       └─ columns: [w_id w_tax]\n" +
			"     └─ Filter\n" +
			"         ├─ AND\n" +
			"         │   ├─ Eq\n" +
			"         │   │   ├─ customer2.c_d_id:1!null\n" +
			"         │   │   └─ 9 (tinyint)\n" +
			"         │   └─ Eq\n" +
			"         │       ├─ customer2.c_id:0!null\n" +
			"         │       └─ 2151 (smallint)\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ keys: [warehouse2.w_id:0!null 9 (tinyint) 2151 (smallint)]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_last c_credit c_discount]\n" +
			"",
	},
	{
		Query: `SELECT d_next_o_id, d_tax FROM district2 WHERE d_w_id = 1 AND d_id = 9 FOR UPDATE`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [district2.d_next_o_id:3, district2.d_tax:2]\n" +
			" └─ IndexedTableAccess(district2)\n" +
			"     ├─ index: [district2.d_w_id,district2.d_id]\n" +
			"     ├─ static: [{[1, 1], [9, 9]}]\n" +
			"     ├─ colSet: (1-11)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: district2\n" +
			"         └─ columns: [d_id d_w_id d_tax d_next_o_id]\n" +
			"",
	},
	{
		Query: `UPDATE district2 SET d_next_o_id = 3002 WHERE d_id = 9 AND d_w_id= 1`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET district2.d_next_o_id:10 = 3002 (smallint))\n" +
			"         └─ IndexedTableAccess(district2)\n" +
			"             ├─ index: [district2.d_w_id,district2.d_id]\n" +
			"             ├─ static: [{[1, 1], [9, 9]}]\n" +
			"             ├─ colSet: (1-11)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: district2\n" +
			"                 └─ columns: [d_id d_w_id d_name d_street_1 d_street_2 d_city d_state d_zip d_tax d_ytd d_next_o_id]\n" +
			"",
	},
	{
		Query: `INSERT INTO orders2 (o_id, o_d_id, o_w_id, o_c_id, o_entry_d, o_ol_cnt, o_all_local) VALUES (3001,9,1,2151,NOW(),12,1)`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Insert(o_id, o_d_id, o_w_id, o_c_id, o_entry_d, o_ol_cnt, o_all_local)\n" +
			"     ├─ InsertDestination\n" +
			"     │   └─ ProcessTable\n" +
			"     │       └─ Table\n" +
			"     │           ├─ name: orders2\n" +
			"     │           └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"     └─ Project\n" +
			"         ├─ columns: [o_id:0!null, o_d_id:1!null, o_w_id:2!null, o_c_id:3, o_entry_d:4, , o_ol_cnt:5, o_all_local:6]\n" +
			"         └─ Values([3001 (smallint),9 (tinyint),1 (tinyint),2151 (smallint),NOW(),12 (tinyint),1 (tinyint)])\n" +
			"",
	},
	{
		Query: `INSERT INTO new_orders2 (no_o_id, no_d_id, no_w_id) VALUES (3001,9,1)`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Insert(no_o_id, no_d_id, no_w_id)\n" +
			"     ├─ InsertDestination\n" +
			"     │   └─ ProcessTable\n" +
			"     │       └─ Table\n" +
			"     │           ├─ name: new_orders2\n" +
			"     │           └─ columns: [no_o_id no_d_id no_w_id]\n" +
			"     └─ Project\n" +
			"         ├─ columns: [no_o_id:0!null, no_d_id:1!null, no_w_id:2!null]\n" +
			"         └─ Values([3001 (smallint),9 (tinyint),1 (tinyint)])\n" +
			"",
	},
	{
		Query: `SELECT i_price, i_name, i_data FROM item2 WHERE i_id = 2532`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [item2.i_price:2, item2.i_name:1, item2.i_data:3]\n" +
			" └─ IndexedTableAccess(item2)\n" +
			"     ├─ index: [item2.i_id]\n" +
			"     ├─ static: [{[2532, 2532]}]\n" +
			"     ├─ colSet: (1-5)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: item2\n" +
			"         └─ columns: [i_id i_name i_price i_data]\n" +
			"",
	},
	{
		Query: `SELECT s_quantity, s_data, s_dist_09 s_dist FROM stock2 WHERE s_i_id = 2532 AND s_w_id= 1 FOR UPDATE`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [stock2.s_quantity:2, stock2.s_data:4, stock2.s_dist_09:3 as s_dist]\n" +
			" └─ IndexedTableAccess(stock2)\n" +
			"     ├─ index: [stock2.s_w_id,stock2.s_i_id]\n" +
			"     ├─ static: [{[1, 1], [2532, 2532]}]\n" +
			"     ├─ colSet: (1-17)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: stock2\n" +
			"         └─ columns: [s_i_id s_w_id s_quantity s_dist_09 s_data]\n" +
			"",
	},
	{
		Query: `UPDATE stock2 SET s_quantity = 39 WHERE s_i_id = 2532 AND s_w_id= 1`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET stock2.s_quantity:2 = 39 (tinyint))\n" +
			"         └─ IndexedTableAccess(stock2)\n" +
			"             ├─ index: [stock2.s_w_id,stock2.s_i_id]\n" +
			"             ├─ static: [{[1, 1], [2532, 2532]}]\n" +
			"             ├─ colSet: (1-17)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: stock2\n" +
			"                 └─ columns: [s_i_id s_w_id s_quantity s_dist_01 s_dist_02 s_dist_03 s_dist_04 s_dist_05 s_dist_06 s_dist_07 s_dist_08 s_dist_09 s_dist_10 s_ytd s_order_cnt s_remote_cnt s_data]\n" +
			"",
	},
	{
		Query: `INSERT INTO order_line2 (ol_o_id, ol_d_id, ol_w_id, ol_number, ol_i_id, ol_supply_w_id, ol_quantity, ol_amount, ol_dist_info) VALUES (
3001,9,1,1,2532,1,5,301,'kkkkkkkkkkkkkkkkkkkkkkkk')`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Insert(ol_o_id, ol_d_id, ol_w_id, ol_number, ol_i_id, ol_supply_w_id, ol_quantity, ol_amount, ol_dist_info)\n" +
			"     ├─ InsertDestination\n" +
			"     │   └─ ProcessTable\n" +
			"     │       └─ Table\n" +
			"     │           ├─ name: order_line2\n" +
			"     │           └─ columns: [ol_o_id ol_d_id ol_w_id ol_number ol_i_id ol_supply_w_id ol_delivery_d ol_quantity ol_amount ol_dist_info]\n" +
			"     └─ Project\n" +
			"         ├─ columns: [ol_o_id:0!null, ol_d_id:1!null, ol_w_id:2!null, ol_number:3!null, ol_i_id:4, ol_supply_w_id:5, , ol_quantity:6, ol_amount:7, ol_dist_info:8]\n" +
			"         └─ Values([3001 (smallint),9 (tinyint),1 (tinyint),1 (tinyint),2532 (smallint),1 (tinyint),5 (tinyint),301 (smallint),kkkkkkkkkkkkkkkkkkkkkkkk (longtext)])\n" +
			"",
	},
	{
		Query: `
-- cycle 1b
SELECT i_price, i_name, i_data FROM item2 WHERE i_id = 2532`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [item2.i_price:2, item2.i_name:1, item2.i_data:3]\n" +
			" └─ IndexedTableAccess(item2)\n" +
			"     ├─ index: [item2.i_id]\n" +
			"     ├─ static: [{[2532, 2532]}]\n" +
			"     ├─ colSet: (1-5)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: item2\n" +
			"         └─ columns: [i_id i_name i_price i_data]\n" +
			"",
	},
	{
		Query: `SELECT s_quantity, s_data, s_dist_09 s_dist FROM stock2 WHERE s_i_id = 2532 AND s_w_id= 1 FOR UPDATE`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [stock2.s_quantity:2, stock2.s_data:4, stock2.s_dist_09:3 as s_dist]\n" +
			" └─ IndexedTableAccess(stock2)\n" +
			"     ├─ index: [stock2.s_w_id,stock2.s_i_id]\n" +
			"     ├─ static: [{[1, 1], [2532, 2532]}]\n" +
			"     ├─ colSet: (1-17)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: stock2\n" +
			"         └─ columns: [s_i_id s_w_id s_quantity s_dist_09 s_data]\n" +
			"",
	},
	{
		Query: `UPDATE stock2 SET s_quantity = 5 WHERE s_i_id = 64568 AND s_w_id= 1`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET stock2.s_quantity:2 = 5 (tinyint))\n" +
			"         └─ IndexedTableAccess(stock2)\n" +
			"             ├─ index: [stock2.s_w_id,stock2.s_i_id]\n" +
			"             ├─ static: [{[1, 1], [64568, 64568]}]\n" +
			"             ├─ colSet: (1-17)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: stock2\n" +
			"                 └─ columns: [s_i_id s_w_id s_quantity s_dist_01 s_dist_02 s_dist_03 s_dist_04 s_dist_05 s_dist_06 s_dist_07 s_dist_08 s_dist_09 s_dist_10 s_ytd s_order_cnt s_remote_cnt s_data]\n" +
			"",
	},
	{
		Query: `INSERT INTO order_line2 (ol_o_id, ol_d_id, ol_w_id, ol_number, ol_i_id, ol_supply_w_id, ol_quantity, ol_amount, ol_dist_info) VALUES (
3001,9,1,11,64568,1,7,298,'oooooooooooooooooooooooo')`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Insert(ol_o_id, ol_d_id, ol_w_id, ol_number, ol_i_id, ol_supply_w_id, ol_quantity, ol_amount, ol_dist_info)\n" +
			"     ├─ InsertDestination\n" +
			"     │   └─ ProcessTable\n" +
			"     │       └─ Table\n" +
			"     │           ├─ name: order_line2\n" +
			"     │           └─ columns: [ol_o_id ol_d_id ol_w_id ol_number ol_i_id ol_supply_w_id ol_delivery_d ol_quantity ol_amount ol_dist_info]\n" +
			"     └─ Project\n" +
			"         ├─ columns: [ol_o_id:0!null, ol_d_id:1!null, ol_w_id:2!null, ol_number:3!null, ol_i_id:4, ol_supply_w_id:5, , ol_quantity:6, ol_amount:7, ol_dist_info:8]\n" +
			"         └─ Values([3001 (smallint),9 (tinyint),1 (tinyint),11 (tinyint),64568 (smallint unsigned),1 (tinyint),7 (tinyint),298 (smallint),oooooooooooooooooooooooo (longtext)])\n" +
			"",
	},
	{
		Query: `
-- cycle 2
UPDATE warehouse2 SET w_ytd = w_ytd + 1767 WHERE w_id = 1`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET warehouse2.w_ytd:8 = (warehouse2.w_ytd:8 + 1767 (smallint)))\n" +
			"         └─ IndexedTableAccess(warehouse2)\n" +
			"             ├─ index: [warehouse2.w_id]\n" +
			"             ├─ static: [{[1, 1]}]\n" +
			"             ├─ colSet: (1-9)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: warehouse2\n" +
			"                 └─ columns: [w_id w_name w_street_1 w_street_2 w_city w_state w_zip w_tax w_ytd]\n" +
			"",
	},
	{
		Query: `SELECT w_street_1, w_street_2, w_city, w_state, w_zip, w_name FROM warehouse2 WHERE w_id = 1`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [warehouse2.w_street_1:2, warehouse2.w_street_2:3, warehouse2.w_city:4, warehouse2.w_state:5, warehouse2.w_zip:6, warehouse2.w_name:1]\n" +
			" └─ IndexedTableAccess(warehouse2)\n" +
			"     ├─ index: [warehouse2.w_id]\n" +
			"     ├─ static: [{[1, 1]}]\n" +
			"     ├─ colSet: (1-9)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: warehouse2\n" +
			"         └─ columns: [w_id w_name w_street_1 w_street_2 w_city w_state w_zip]\n" +
			"",
	},
	{
		Query: `UPDATE district2 SET d_ytd = d_ytd + 1767 WHERE d_w_id = 1 AND d_id= 8`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET district2.d_ytd:9 = (district2.d_ytd:9 + 1767 (smallint)))\n" +
			"         └─ IndexedTableAccess(district2)\n" +
			"             ├─ index: [district2.d_w_id,district2.d_id]\n" +
			"             ├─ static: [{[1, 1], [8, 8]}]\n" +
			"             ├─ colSet: (1-11)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: district2\n" +
			"                 └─ columns: [d_id d_w_id d_name d_street_1 d_street_2 d_city d_state d_zip d_tax d_ytd d_next_o_id]\n" +
			"",
	},
	{
		Query: `SELECT d_street_1, d_street_2, d_city, d_state, d_zip, d_name FROM district2 WHERE d_w_id = 1 AND d_id = 8`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [district2.d_street_1:3, district2.d_street_2:4, district2.d_city:5, district2.d_state:6, district2.d_zip:7, district2.d_name:2]\n" +
			" └─ IndexedTableAccess(district2)\n" +
			"     ├─ index: [district2.d_w_id,district2.d_id]\n" +
			"     ├─ static: [{[1, 1], [8, 8]}]\n" +
			"     ├─ colSet: (1-11)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: district2\n" +
			"         └─ columns: [d_id d_w_id d_name d_street_1 d_street_2 d_city d_state d_zip]\n" +
			"",
	},
	{
		Query: `SELECT count(c_id) namecnt FROM customer2 WHERE c_w_id = 1 AND c_d_id= 5 AND c_last='ESEEINGABLE'`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [count(customer2.c_id):0!null as namecnt]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNT(customer2.c_id:0!null)\n" +
			"     ├─ group: \n" +
			"     └─ Filter\n" +
			"         ├─ Eq\n" +
			"         │   ├─ customer2.c_last:3\n" +
			"         │   └─ ESEEINGABLE (longtext)\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ static: [{[1, 1], [5, 5], [NULL, ∞)}]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_last]\n" +
			"",
	},
	{
		Query: `SELECT c_id FROM customer2 WHERE c_w_id = 1 AND c_d_id= 5 AND c_last='ESEEINGABLE' ORDER BY c_first`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [customer2.c_id:0!null]\n" +
			" └─ Sort(customer2.c_first:3 ASC nullsFirst)\n" +
			"     └─ Filter\n" +
			"         ├─ Eq\n" +
			"         │   ├─ customer2.c_last:5\n" +
			"         │   └─ ESEEINGABLE (longtext)\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ static: [{[1, 1], [5, 5], [NULL, ∞)}]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_first c_middle c_last c_street_1 c_street_2 c_city c_state c_zip c_phone c_since c_credit c_credit_lim c_discount c_balance c_ytd_payment c_payment_cnt c_delivery_cnt c_data]\n" +
			"",
	},
	{
		Query: `SELECT c_first, c_middle, c_last, c_street_1, c_street_2, c_city, c_state, c_zip, c_phone, c_credit, c_credit_lim, c_discount, c_balance, c_ytd_payment, c_since FROM customer2 WHERE c_w_id = 1 AND c_d_id= 5 AND c_id=1838 FOR UPDATE`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [customer2.c_first:3, customer2.c_middle:4, customer2.c_last:5, customer2.c_street_1:6, customer2.c_street_2:7, customer2.c_city:8, customer2.c_state:9, customer2.c_zip:10, customer2.c_phone:11, customer2.c_credit:13, customer2.c_credit_lim:14, customer2.c_discount:15, customer2.c_balance:16, customer2.c_ytd_payment:17, customer2.c_since:12]\n" +
			" └─ IndexedTableAccess(customer2)\n" +
			"     ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"     ├─ static: [{[1, 1], [5, 5], [1838, 1838]}]\n" +
			"     ├─ colSet: (1-21)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: customer2\n" +
			"         └─ columns: [c_id c_d_id c_w_id c_first c_middle c_last c_street_1 c_street_2 c_city c_state c_zip c_phone c_since c_credit c_credit_lim c_discount c_balance c_ytd_payment]\n" +
			"",
	},
	{
		Query: `UPDATE customer2 SET c_balance=-1777.000000, c_ytd_payment=1777.000000 WHERE c_w_id = 1 AND c_d_id=5 AND c_id=1838`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Update\n" +
			"     └─ UpdateSource(SET customer2.c_balance:16 = -1777.000000,SET customer2.c_ytd_payment:17 = 1777 (decimal(10,6)))\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ static: [{[1, 1], [5, 5], [1838, 1838]}]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_first c_middle c_last c_street_1 c_street_2 c_city c_state c_zip c_phone c_since c_credit c_credit_lim c_discount c_balance c_ytd_payment c_payment_cnt c_delivery_cnt c_data]\n" +
			"",
	},
	{
		Query: `INSERT INTO history2 (h_c_d_id, h_c_w_id, h_c_id, h_d_id, h_w_id, h_date, h_amount, h_data) VALUES (5,1,1838,8,1,NOW(),1767,'name-rqojn name-dnvgs ')`,
		ExpectedPlan: "RowUpdateAccumulator\n" +
			" └─ Insert(h_c_d_id, h_c_w_id, h_c_id, h_d_id, h_w_id, h_date, h_amount, h_data)\n" +
			"     ├─ InsertDestination\n" +
			"     │   └─ ProcessTable\n" +
			"     │       └─ Table\n" +
			"     │           ├─ name: history2\n" +
			"     │           └─ columns: [h_c_id h_c_d_id h_c_w_id h_d_id h_w_id h_date h_amount h_data]\n" +
			"     └─ Project\n" +
			"         ├─ columns: [h_c_id:2, h_c_d_id:0, h_c_w_id:1, h_d_id:3, h_w_id:4, h_date:5, h_amount:6, h_data:7]\n" +
			"         └─ Values([5 (tinyint),1 (tinyint),1838 (smallint),8 (tinyint),1 (tinyint),NOW(),1767 (smallint),name-rqojn name-dnvgs  (longtext)])\n" +
			"",
	},
	{
		Query: `
-- cycle 3
SELECT count(c_id) namecnt FROM customer2 WHERE c_w_id = 1 AND c_d_id= 1 AND c_last='PRIESEPRES'`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [count(customer2.c_id):0!null as namecnt]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNT(customer2.c_id:0!null)\n" +
			"     ├─ group: \n" +
			"     └─ Filter\n" +
			"         ├─ Eq\n" +
			"         │   ├─ customer2.c_last:3\n" +
			"         │   └─ PRIESEPRES (longtext)\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ static: [{[1, 1], [1, 1], [NULL, ∞)}]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_last]\n" +
			"",
	},
	{
		Query: `SELECT c_balance, c_first, c_middle, c_id FROM customer2 WHERE c_w_id = 1 AND c_d_id= 1 AND c_last='PRIESEPRES' ORDER BY c_first`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [customer2.c_balance:16, customer2.c_first:3, customer2.c_middle:4, customer2.c_id:0!null]\n" +
			" └─ Sort(customer2.c_first:3 ASC nullsFirst)\n" +
			"     └─ Filter\n" +
			"         ├─ Eq\n" +
			"         │   ├─ customer2.c_last:5\n" +
			"         │   └─ PRIESEPRES (longtext)\n" +
			"         └─ IndexedTableAccess(customer2)\n" +
			"             ├─ index: [customer2.c_w_id,customer2.c_d_id,customer2.c_id]\n" +
			"             ├─ static: [{[1, 1], [1, 1], [NULL, ∞)}]\n" +
			"             ├─ colSet: (1-21)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: customer2\n" +
			"                 └─ columns: [c_id c_d_id c_w_id c_first c_middle c_last c_street_1 c_street_2 c_city c_state c_zip c_phone c_since c_credit c_credit_lim c_discount c_balance c_ytd_payment c_payment_cnt c_delivery_cnt c_data]\n" +
			"",
	},
	{
		Query: `SELECT o_id, o_carrier_id, o_entry_d FROM orders2 WHERE o_w_id = 1 AND o_d_id = 1 AND o_c_id = 355 ORDER BY o_id DESC`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [orders2.o_id:0!null, orders2.o_carrier_id:5, orders2.o_entry_d:4]\n" +
			" └─ Sort(orders2.o_id:0!null DESC nullsFirst)\n" +
			"     └─ Filter\n" +
			"         ├─ Eq\n" +
			"         │   ├─ orders2.o_c_id:3\n" +
			"         │   └─ 355 (smallint)\n" +
			"         └─ IndexedTableAccess(orders2)\n" +
			"             ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"             ├─ static: [{[1, 1], [1, 1], [NULL, ∞)}]\n" +
			"             ├─ colSet: (1-8)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: orders2\n" +
			"                 └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"",
	},
	{
		Query: `SELECT ol_i_id, ol_supply_w_id, ol_quantity, ol_amount, ol_delivery_d FROM order_line2 WHERE ol_w_id = 1 AND ol_d_id = 1 AND ol_o_id = 1`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [order_line2.ol_i_id:3, order_line2.ol_supply_w_id:4, order_line2.ol_quantity:6, order_line2.ol_amount:7, order_line2.ol_delivery_d:5]\n" +
			" └─ IndexedTableAccess(order_line2)\n" +
			"     ├─ index: [order_line2.ol_w_id,order_line2.ol_d_id,order_line2.ol_o_id,order_line2.ol_number]\n" +
			"     ├─ static: [{[1, 1], [1, 1], [1, 1], [NULL, ∞)}]\n" +
			"     ├─ colSet: (1-10)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: order_line2\n" +
			"         └─ columns: [ol_o_id ol_d_id ol_w_id ol_i_id ol_supply_w_id ol_delivery_d ol_quantity ol_amount]\n" +
			"",
	},
	{
		Query: `
-- cycle 4
SELECT d_next_o_id FROM district2 WHERE d_id = 5 AND d_w_id= 1`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [district2.d_next_o_id:2]\n" +
			" └─ IndexedTableAccess(district2)\n" +
			"     ├─ index: [district2.d_w_id,district2.d_id]\n" +
			"     ├─ static: [{[1, 1], [5, 5]}]\n" +
			"     ├─ colSet: (1-11)\n" +
			"     ├─ tableId: 1\n" +
			"     └─ Table\n" +
			"         ├─ name: district2\n" +
			"         └─ columns: [d_id d_w_id d_next_o_id]\n" +
			"",
	},
	{
		Query: `SELECT COUNT(DISTINCT (s_i_id)) FROM order_line2, stock2 WHERE ol_w_id = 1 AND ol_d_id = 5 AND ol_o_id < 3003 AND ol_o_id >= 2983 AND s_w_id= 1 AND s_i_id=ol_i_id AND s_quantity < 18`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [countdistinct([stock2.s_i_id]):0!null as COUNT(DISTINCT (s_i_id))]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNTDISTINCT([stock2.s_i_id])\n" +
			"     ├─ group: \n" +
			"     └─ LookupJoin\n" +
			"         ├─ IndexedTableAccess(order_line2)\n" +
			"         │   ├─ index: [order_line2.ol_w_id,order_line2.ol_d_id,order_line2.ol_o_id,order_line2.ol_number]\n" +
			"         │   ├─ static: [{[1, 1], [5, 5], [2983, 3003), [NULL, ∞)}]\n" +
			"         │   ├─ colSet: (1-10)\n" +
			"         │   ├─ tableId: 1\n" +
			"         │   └─ Table\n" +
			"         │       ├─ name: order_line2\n" +
			"         │       └─ columns: [ol_o_id ol_d_id ol_w_id ol_i_id]\n" +
			"         └─ Filter\n" +
			"             ├─ AND\n" +
			"             │   ├─ Eq\n" +
			"             │   │   ├─ stock2.s_w_id:1!null\n" +
			"             │   │   └─ 1 (tinyint)\n" +
			"             │   └─ LessThan\n" +
			"             │       ├─ stock2.s_quantity:2\n" +
			"             │       └─ 18 (tinyint)\n" +
			"             └─ IndexedTableAccess(stock2)\n" +
			"                 ├─ index: [stock2.s_w_id,stock2.s_i_id]\n" +
			"                 ├─ keys: [1 (tinyint) order_line2.ol_i_id:3]\n" +
			"                 ├─ colSet: (11-27)\n" +
			"                 ├─ tableId: 2\n" +
			"                 └─ Table\n" +
			"                     ├─ name: stock2\n" +
			"                     └─ columns: [s_i_id s_w_id s_quantity]\n" +
			"",
	},
	{
		Query: `
-- other
SELECT o_id, o_entry_d, COALESCE(o_carrier_id,0)
FROM orders2
WHERE
  o_w_id = 1 AND
  o_d_id = 3 AND
  o_c_id = 20001 AND
  o_id = (SELECT MAX(o_id) FROM orders2 WHERE o_w_id = 1 AND o_d_id = 3 AND o_c_id = 20001)`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [orders2.o_id:0!null, orders2.o_entry_d:4, coalesce(orders2.o_carrier_id:5,0 (tinyint)) as COALESCE(o_carrier_id,0)]\n" +
			" └─ Filter\n" +
			"     ├─ AND\n" +
			"     │   ├─ Eq\n" +
			"     │   │   ├─ orders2.o_id:0!null\n" +
			"     │   │   └─ Subquery\n" +
			"     │   │       ├─ cacheable: true\n" +
			"     │   │       ├─ alias-string: select MAX(o_id) from orders2 where o_w_id = 1 and o_d_id = 3 and o_c_id = 20001\n" +
			"     │   │       └─ Project\n" +
			"     │   │           ├─ columns: [max(orders2.o_id):8!null as MAX(o_id)]\n" +
			"     │   │           └─ GroupBy\n" +
			"     │   │               ├─ select: MAX(orders2.o_id:8!null)\n" +
			"     │   │               ├─ group: \n" +
			"     │   │               └─ Filter\n" +
			"     │   │                   ├─ Eq\n" +
			"     │   │                   │   ├─ orders2.o_c_id:11\n" +
			"     │   │                   │   └─ 20001 (smallint)\n" +
			"     │   │                   └─ IndexedTableAccess(orders2)\n" +
			"     │   │                       ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"     │   │                       ├─ static: [{[1, 1], [3, 3], [NULL, ∞)}]\n" +
			"     │   │                       ├─ colSet: (9-16)\n" +
			"     │   │                       ├─ tableId: 2\n" +
			"     │   │                       └─ Table\n" +
			"     │   │                           ├─ name: orders2\n" +
			"     │   │                           └─ columns: [o_id o_d_id o_w_id o_c_id]\n" +
			"     │   └─ Eq\n" +
			"     │       ├─ orders2.o_c_id:3\n" +
			"     │       └─ 20001 (smallint)\n" +
			"     └─ IndexedTableAccess(orders2)\n" +
			"         ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"         ├─ static: [{[1, 1], [3, 3], [NULL, ∞)}]\n" +
			"         ├─ colSet: (1-8)\n" +
			"         ├─ tableId: 1\n" +
			"         └─ Table\n" +
			"             ├─ name: orders2\n" +
			"             └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"",
	},
	{
		Query: `
select o_id, o.o_d_id
from
  orders2 o,
  (
    select o_c_id, o_w_id, o_d_id, count(distinct o_id)
    from orders2
    where o_w_id=1  and o_id > 2100 and o_id < 11153
    group by o_c_id,o_d_id,o_w_id
    having count( distinct o_id) > 1
    limit 1
  ) t
  where
    t.o_w_id=o.o_w_id and
    t.o_d_id=o.o_d_id and
    t.o_c_id=o.o_c_id
  limit 1;`,
		ExpectedPlan: "Limit(1)\n" +
			" └─ Project\n" +
			"     ├─ columns: [o.o_id:4!null, o.o_d_id:5!null]\n" +
			"     └─ HashJoin\n" +
			"         ├─ AND\n" +
			"         │   ├─ AND\n" +
			"         │   │   ├─ Eq\n" +
			"         │   │   │   ├─ t.o_w_id:1!null\n" +
			"         │   │   │   └─ o.o_w_id:6!null\n" +
			"         │   │   └─ Eq\n" +
			"         │   │       ├─ t.o_d_id:2!null\n" +
			"         │   │       └─ o.o_d_id:5!null\n" +
			"         │   └─ Eq\n" +
			"         │       ├─ t.o_c_id:0\n" +
			"         │       └─ o.o_c_id:7\n" +
			"         ├─ SubqueryAlias\n" +
			"         │   ├─ name: t\n" +
			"         │   ├─ outerVisibility: false\n" +
			"         │   ├─ isLateral: false\n" +
			"         │   ├─ cacheable: true\n" +
			"         │   ├─ colSet: (19-22)\n" +
			"         │   ├─ tableId: 3\n" +
			"         │   └─ Limit(1)\n" +
			"         │       └─ Project\n" +
			"         │           ├─ columns: [orders2.o_c_id:1, orders2.o_w_id:2!null, orders2.o_d_id:3!null, countdistinct([orders2.o_id]):0!null as count(distinct o_id)]\n" +
			"         │           └─ Having\n" +
			"         │               ├─ GreaterThan\n" +
			"         │               │   ├─ countdistinct([orders2.o_id]):0!null\n" +
			"         │               │   └─ 1 (tinyint)\n" +
			"         │               └─ GroupBy\n" +
			"         │                   ├─ select: COUNTDISTINCT([orders2.o_id]), orders2.o_c_id:3, orders2.o_w_id:2!null, orders2.o_d_id:1!null, orders2.o_id:0!null\n" +
			"         │                   ├─ group: orders2.o_c_id:3, orders2.o_d_id:1!null, orders2.o_w_id:2!null\n" +
			"         │                   └─ IndexedTableAccess(orders2)\n" +
			"         │                       ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"         │                       ├─ static: [{[1, 1], [NULL, ∞), (2100, 11153)}]\n" +
			"         │                       ├─ colSet: (9-16)\n" +
			"         │                       ├─ tableId: 2\n" +
			"         │                       └─ Table\n" +
			"         │                           ├─ name: orders2\n" +
			"         │                           └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"         └─ HashLookup\n" +
			"             ├─ left-key: TUPLE(t.o_w_id:1!null, t.o_d_id:2!null, t.o_c_id:0)\n" +
			"             ├─ right-key: TUPLE(o.o_w_id:2!null, o.o_d_id:1!null, o.o_c_id:3)\n" +
			"             └─ TableAlias(o)\n" +
			"                 └─ ProcessTable\n" +
			"                     └─ Table\n" +
			"                         ├─ name: orders2\n" +
			"                         └─ columns: [o_id o_d_id o_w_id o_c_id]\n" +
			"",
	},
}