# Band-Protocol-Assignment-1 Boss Baby's Revenge

## Overview
ในโจทย์ข้อที่ 1 นี้ เป็นการแก้ปัญหาให้กับ Boss baby บอสเบบี้จะยิงกลับอย่างน้อยหนึ่งครั้ง แต่เฉพาะเมื่อโดนยิงมาก่อนเท่านั้น โดยเรา ช่วยตรวจสอบว่าบอสเบบี้ได้แก้แค้นทุกครั้งที่ถูกยิงและไม่เคยเริ่มยิงก่อน โดยจะรับ string เข้ามา 1 ชุดเป็นตัวแทนเหตุการณ์ ซึ่ง string จะประกอบไปด้วย 'S' และ 'R' โดยที่ 'S' หมายถึงการยิงและ 'R' หมายถึงการแก้แค้น และคืนค่า "Good boy" ถ้าทุกการยิงถูกแก้แค้นอย่างน้อยหนึ่งครั้งและบอสเบบี้ไม่เคยเริ่มยิงก่อน คืนค่า "Bad boy" ถ้าเงื่อนไขเหล่านี้ไม่เป็นจริง บอสเบบี้ไม่จำเป็นต้องยิงกลับก่อนการยิงครั้งถัดไปของเด็กๆ เขาสามารถยิงกลับเมื่อไรก็ได้ตราบเท่าที่เขาไม่เริ่มยิงก่อน

## Explain logic
- ทำการรับ input และแปลงให้อยู่ในรูป array
- เช็คว่า array[0] ไม่เป็น 'R' และ array[last] ไม่เป็น 'S' 
- วนลูป เช็คว่า หากมี 'S' ให้ค่าการโดนยิงเพิ่ม 1 หากตรวจพบ 'R' ให้ค่าการโดนยิงลบ 1
- หลังจากวนลูปเสร็จให้เช็คว่าถ้าค่าการโดนยิง(shorts) <= 0 แสดงว่าได้รับการแก้แค้นทั้งหมด คืนค่า "Good boy" แต่ถ้าไม่ คืนค่า "Bad boy"

## Requirements
- Go 1.21.5+

## Getting Started
### Installation and Running the Project
1. Clone the project from GitHub:
   ```sh
   git clone https://github.com/PParist/band-protocol-assignment-1.git
2. Installation Go:
   ```sh
   Download and install https://go.dev/doc/install
3. Running the Project:
   ```sh
   cd assignment_1
   go run main.go / go run .
